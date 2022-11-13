package tarsserver

import (
	"context"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
)

type ConnectHander struct {
	svrImp    *TarsServer
	conn      net.Conn
	sendLock  *sync.Mutex
	isClosed  bool
	idleTime  time.Time
	invokeNum int32
	sendchan  chan []byte
	closechan chan struct{}
	closeOnce *sync.Once
}

func NewConnectHandler(svrImp *TarsServer, conn net.Conn) *ConnectHander {
	ch := &ConnectHander{
		svrImp:    svrImp,
		conn:      conn,
		isClosed:  false,
		sendLock:  &sync.Mutex{},
		closeOnce: &sync.Once{},
		sendchan:  make(chan []byte, 10000),
		closechan: make(chan struct{}),
	}
	return ch
}

func (ch *ConnectHander) recv() {
	defer func() {
		<-ch.svrImp.acceptCounter
		ch.close("recv defer")
	}()
	cfg := ch.svrImp.conf
	conn := ch.conn
	buffer := make([]byte, 1024*4)
	if cfg.Proto == "udp" {
		buffer = make([]byte, 1024*1024*10)
	}
	var currBuffer []byte
	ch.idleTime = time.Now()
	var n int
	var err error
	var udpAddr *net.UDPAddr
	for {
		if ch.isClosed {
			return
		}
		conn.SetReadDeadline(time.Now().Add(cfg.ReadTimeout))
		if cfg.Proto == "udp" {
			n, udpAddr, err = (conn.(*net.UDPConn)).ReadFromUDP(buffer)
		} else {
			n, err = conn.Read(buffer)
		}
		if err != nil {
			if cfg.Proto == "tcp" &&
				len(currBuffer) == 0 &&
				ch.invokeNum == 0 &&
				ch.idleTime.Add(cfg.IdleTimeout).Before(time.Now()) {
				appzaplog.Info("close idle connection", zap.Any("localaddr", conn.LocalAddr()), zap.Any("remoteaddr", conn.RemoteAddr()))
				return
			}
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue // no data, not error
			}
			if err == io.EOF {
				appzaplog.Warn("connection closed by remote",
					zap.Any("localaddr", conn.LocalAddr()),
					zap.Any("remoteaddr", conn.RemoteAddr()))
			} else {
				appzaplog.Error("read packge error", zap.Error(err))
			}
			return
		}
		ch.idleTime = time.Now()
		currBuffer = append(currBuffer, buffer[:n]...)
		for {
			pkgLen, status := ch.svrImp.protocolImp.ParsePackage(currBuffer)
			if status == PACKAGE_LESS {
				break
			}
			if status == PACKAGE_FULL {
				pkg := make([]byte, pkgLen-4)
				copy(pkg, currBuffer[4:pkgLen])
				currBuffer = currBuffer[pkgLen:]

				if udpAddr != nil {
					go ch.invokeUdp(pkg, udpAddr)
				} else {
					go ch.invoke(pkg)
				}

				if len(currBuffer) > 0 {
					continue
				} else if len(currBuffer) == 0 {
					//TODO may not free
					currBuffer = nil
				}
				break
			}
			// status error
			if cfg.Proto == "tcp" {
				appzaplog.Error("parse package error")
				return
			}
		}
	}
}

func (ch *ConnectHander) sendUdp(rsp []byte, udpAddr *net.UDPAddr) (err error) {
	ch.sendLock.Lock()
	defer ch.sendLock.Unlock()
	cfg := ch.svrImp.conf
	ch.conn.SetWriteDeadline(time.Now().Add(cfg.WriteTimeout))
	if udpAddr != nil {
		_, err = ch.conn.(*net.UDPConn).WriteToUDP(rsp, udpAddr)
	} else {
		return errors.New("udpAddr nil")
	}
	return
}

func (ch *ConnectHander) sendTcp(rsp []byte) (err error) {
	// catch the panic for a close send chan
	defer func() {
		if perr := recover(); perr != nil {
			appzaplog.Error("sendTcp panic", zap.Any("recover", perr))
		}
	}()
	if ch.isClosed {
		return nil
	}
	select {
	case <-ch.closechan:
		return errors.New("TCP closed")
	case ch.sendchan <- rsp:
		return nil
	default:
		return errors.New("TCP send buffer full")
	}
}

func (ch *ConnectHander) loopsend() {
	defer ch.close("loopsend defer")
	for {
		if ch.isClosed {
			return
		}
		select {
		case resp := <-ch.sendchan:
			//ch.conn.SetWriteDeadline(time.Now().Add(ch.svrImp.conf.WriteTimeout))
			{
				_, err := ch.conn.Write(resp)
				if err != nil {
					appzaplog.Error("send error", zap.Error(err))
					return
				}
			}
		case <-ch.closechan:
			return
		}
	}
}

// only process tcp
func (ch *ConnectHander) invoke(pkg []byte) {
	select {
	case ch.svrImp.invokeCounter <- true:
		atomic.AddInt32(&ch.invokeNum, 1)
	default:
		appzaplog.Warn("drop invoke since invokequeue is full, need larger queue size")
		return
	}
	defer func() {
		atomic.AddInt32(&ch.invokeNum, -1)
		<-ch.svrImp.invokeCounter
	}()

	ch.svrImp.lastInvoke = time.Now()
	cfg := ch.svrImp.conf
	rspChan := make(chan []byte, 1)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithTimeout(context.Background(), cfg.HandleTimeout)
	defer cancel()

	go func() {
		rsp, err := ch.svrImp.protocolImp.Invoke(ctx, pkg)
		if err != nil {
			appzaplog.Error("invoke error", zap.Error(err))
			rspChan <- nil
			return
		}
		rspChan <- rsp
	}()
	select {
	case <-ctx.Done():
		rsp, err := ch.svrImp.protocolImp.InvokeTimeout(pkg)
		if err != nil {
			appzaplog.Error("invoke timeout error", zap.Error(err))
			ch.close("invoke timeout error")
			return
		}
		ch.sendTcp(rsp)
	case rsp := <-rspChan:
		if rsp == nil {
			ch.close("rsp is nil")
			return
		}
		ch.sendTcp(rsp)
	}
}

func (ch *ConnectHander) close(reason string) {
	ch.closeOnce.Do(func() {
		ch.isClosed = true
		appzaplog.Info("connection close", zap.String("reason", reason),
			zap.Any("localaddr", ch.conn.LocalAddr()),
			zap.Any("remoteaddr", ch.conn.RemoteAddr()))
		ch.conn.Close()
		close(ch.closechan)
	})
}

func (ch *ConnectHander) invokeUdp(pkg []byte, udpAddr *net.UDPAddr) {
	atomic.AddInt32(&ch.invokeNum, 1)
	ch.svrImp.invokeCounter <- true
	defer func() {
		atomic.AddInt32(&ch.invokeNum, -1)
		<-ch.svrImp.invokeCounter
	}()
	ch.svrImp.lastInvoke = time.Now()
	cfg := ch.svrImp.conf
	rspChan := make(chan []byte)
	go func() {
		rsp, err := ch.svrImp.protocolImp.Invoke(context.TODO(), pkg)
		if err != nil {
			appzaplog.Error("invoke error", zap.Error(err))
			//fmt.Println("invoke error", err)
			rspChan <- nil
			return
		}
		rspChan <- rsp
	}()
	select {
	case <-time.After(cfg.HandleTimeout):
		rsp, err := ch.svrImp.protocolImp.InvokeTimeout(pkg)
		if err != nil {
			appzaplog.Error("invoke timeout error", zap.Error(err))
			ch.close("invoke timeout error")
			return
		}
		ch.sendUdp(rsp, udpAddr)
	case rsp := <-rspChan:
		if rsp == nil {
			ch.close("rsp is nil")
			return
		}
		ch.sendUdp(rsp, udpAddr)
	}
}
