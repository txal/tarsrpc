// @author kordenlu
// @创建时间 2018/02/11 19:41
// 功能描述:

package tarsserver

import (
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
)

type connection struct {
	tc *TarsClient

	idleTime  time.Time
	invokeNum int32

	connLock *sync.Mutex // protect blow fields
	conn     net.Conn
	isClosed bool
}

func (c *connection) send(conn net.Conn, connDone chan bool) {
	var req []byte
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-connDone: // connection closed
			return
		default:
			select {
			case req = <-c.tc.sendQueue: // Fetch jobs
			case <-t.C:
				if c.isClosed {
					return
				}
				// TODO: check one-way invoke for idle detect
				if c.invokeNum == 0 && c.idleTime.Add(c.tc.conf.IdleTimeout).Before(time.Now()) {
					appzaplog.Info("close idle connection",
						zap.Any("localaddr", conn.LocalAddr()),
						zap.Any("remoteaddr", conn.RemoteAddr()))
					c.close(conn)
					return
				}
				continue
			}
		}
		atomic.AddInt32(&c.invokeNum, 1)
		if c.tc.conf.WriteTimeout != 0 {
			conn.SetWriteDeadline(time.Now().Add(c.tc.conf.WriteTimeout))
		}
		c.idleTime = time.Now()
		_, err := conn.Write(req)
		if err != nil {
			//TODO add retry time
			c.tc.sendQueue <- req
			appzaplog.Error("send request error:", zap.Error(err),
				zap.Any("localaddr", conn.LocalAddr()),
				zap.Any("remoteaddr", conn.RemoteAddr()))
			c.close(conn)
			return
		}
	}
}

func (c *connection) recv(conn net.Conn, connDone chan bool) {
	defer func() {
		connDone <- true
	}()
	buffer := make([]byte, 1024*4)
	var currBuffer []byte
	var n int
	var err error
	for {
		conn.SetReadDeadline(time.Now().Add(c.tc.conf.ReadTimeout))
		n, err = conn.Read(buffer)
		if err != nil {
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue // no data, not error
			}
			if opErr, ok := err.(*net.OpError); ok {
				appzaplog.Warn("net op error", zap.Any("error", opErr.Error()),
					zap.Any("localaddr", conn.LocalAddr()),
					zap.Any("remoteaddr", conn.RemoteAddr()))
				c.close(conn)
				return // connection is closed
			}
			if err == io.EOF {
				appzaplog.Info("connection closed by remote",
					zap.Any("localaddr", conn.LocalAddr()),
					zap.Any("remoteaddr", conn.RemoteAddr()))
			} else {
				appzaplog.Warn("read package error", zap.Error(err),
					zap.Any("localaddr", conn.LocalAddr()),
					zap.Any("remoteaddr", conn.RemoteAddr()))
			}
			c.close(conn)
			return
		}
		currBuffer = append(currBuffer, buffer[:n]...)
		for {
			pkgLen, status := c.tc.cp.ParsePackage(currBuffer)
			if status == PACKAGE_LESS {
				break
			}
			if status == PACKAGE_FULL {
				atomic.AddInt32(&c.invokeNum, -1)
				pkg := make([]byte, pkgLen-4)
				copy(pkg, currBuffer[4:pkgLen])
				currBuffer = currBuffer[pkgLen:]
				go c.tc.cp.Recv(pkg)
				if len(currBuffer) > 0 {
					continue
				}
				currBuffer = nil
				break
			}
			appzaplog.Error("parse package error",
				zap.Any("localaddr", conn.LocalAddr()),
				zap.Any("remoteaddr", conn.RemoteAddr()))
			c.close(conn)
			return
		}
	}
}

var (
	NetDialTimeoutErr = errors.New("netDialTimeout")
)

func (c *connection) reConnect() (err error) {
	if c.isClosed {
		c.connLock.Lock()
		if c.isClosed {
			appzaplog.Debug("ReConnect", zap.Any("addr", c.tc.address))
			// todo make timeout configable
			c.conn, err = net.DialTimeout(c.tc.conf.Proto, c.tc.address, 1*time.Second)
			if err != nil {
				c.connLock.Unlock()
				return NetDialTimeoutErr
			}
			appzaplog.Info("ReConnect",
				zap.Any("localaddr", c.conn.LocalAddr()),
				zap.Any("remoteaddr", c.conn.RemoteAddr()))
			c.idleTime = time.Now()
			c.isClosed = false
			connDone := make(chan bool, 1)
			go c.recv(c.conn, connDone)
			go c.send(c.conn, connDone)
		}
		c.connLock.Unlock()
	}
	return nil
}

func (c *connection) close(conn net.Conn) {
	c.connLock.Lock()
	if c != nil && conn == c.conn {
		c.isClosed = true
		if conn != nil {
			conn.Close()
		}
	}
	c.connLock.Unlock()
}
