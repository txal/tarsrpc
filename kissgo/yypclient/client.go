package yypclient

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"code.com/tars/goframework/kissgo/s2s"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"gopkg.in/fatih/pool.v2"
	"code.com/tars/goframework/kissgo/yyp"
)

type addr struct {
	network string
	str     string
}

// Client is a abstraction of connections to YYP servers.
type Client struct {
	// S2SName is the target server name, registered in s2s.
	S2SName string
	// Groups are the prioritized selected groups of target servers. Primary group comes first.
	Groups []int
	// Timeout in millisecond
	Timeout int
	// Address is used to bypass s2s
	Address []addr

	l               *log.Logger
	addressNotifier <-chan s2s.S2sMetaVector

	addressBook       map[int64]map[int]map[int]addr // Server -> Group -> ISPType -> address
	addressLock       sync.RWMutex
	groupToServerList map[int][]int64     // Group -> List<Server>
	serverToPool      map[int64]pool.Pool // Server -> TCPConnectionPool
}

func init() {
	err := s2s.InitS2s("HanMeiMei", "2613ef093eb54045bb58f16ba60edaa3d8e25a913019820b")
	if err != nil {
		log.Fatalf("init s2s fail: %v", err)
	}
	rand.Seed(time.Now().Unix())
}

// Setup initialize the yyp client
// It must be called only once for every client, before calling any other method.
func (t *Client) Setup() {
	if t.l == nil {
		t.l = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	}
	if t.addressNotifier == nil {
		if t.S2SName == "" {
			t.l.Panicf("S2SName is empty")
		}
		if len(t.Groups) == 0 {
			t.l.Panicf("Groups is empty")
		}
		ch, err := s2s.Subscribe(t.S2SName, 0)
		if err != nil {
			t.l.Printf("Error s2s.Subscribe(%s): %v", t.S2SName, err)
			return
		}
		t.addressNotifier = ch
		go t.updateAddress()
		t.l.Printf("Setup done with S2SName: %q", t.S2SName)
	}
	if t.serverToPool == nil {
		t.serverToPool = make(map[int64]pool.Pool)
	}
	if t.Address != nil {
		t.groupToServerList = make(map[int][]int64)
		g := t.Groups[0]
		for i, a := range t.Address {
			t.groupToServerList[g] = []int64{int64(i)}
			p, _ := pool.NewChannelPool(0, 6000, func() (net.Conn, error) {
				return net.Dial(a.network, a.str)
			})
			t.serverToPool[int64(i)] = p
		}
	}
}

// DoRoundtripRetry calls DoRoundtrip with a number of retries.
func (t *Client) DoRoundtripRetry(reqURI, rspURI int, req []byte) (rsp []byte, err error) {
	for i := 0; i < 3; i++ {
		rsp, err = t.DoRoundtrip(reqURI, rspURI, req)
		if err != nil {
			t.l.Printf("%v", err)
			continue
		}
		break
	}
	return
}

type YYPResp struct {
	Uri uint32
	Reserve uint16
	Data []byte
}

// DoRoundtrip send something on this session, and wait until recv something.
// The caller should retry when error happens.
func (t *Client) DoRoundtrip(reqURI, rspURI int, req []byte) (rsp []byte, err error) {
	conn, err := t.getConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	t.l.Printf("yyp.Client.DoRoundtrip to %q : %v", t.S2SName, conn.RemoteAddr())
	err = conn.SetWriteDeadline(time.Time{})
	if err != nil {
		markUnusable(conn)
		return nil, fmt.Errorf("yyp.Client.DoRoundtrip SetWriteDeadline: %v", err)
	}
	if t.Timeout > 0 {
		err = conn.SetReadDeadline(time.Now().Add(time.Duration(t.Timeout) * time.Millisecond))
	} else {
		err = conn.SetReadDeadline(time.Time{})
	}
	if err != nil {
		markUnusable(conn)
		return nil, fmt.Errorf("yyp.Client.DoRoundtrip SetReadDeadline: %v", err)
	}

	var buf bytes.Buffer
	var scratch [64]byte
	var b []byte

	b = scratch[0:4]
	binary.LittleEndian.PutUint32(b, uint32(len(req)+10))
	buf.Write(b)

	b = scratch[0:4]
	binary.LittleEndian.PutUint32(b, uint32(reqURI))
	buf.Write(b)

	b = scratch[0:2]
	binary.LittleEndian.PutUint16(b, 200)
	buf.Write(b)

	buf.Write(req)
	bb := buf.Bytes()
	//t.l.Printf("yyp.Client.DoRoundtrip to %q : write buf len = %d", t.S2SName, len(bb))
	//fmt.Printf(hex.Dump(bb))

	_, err = conn.Write(bb)
	if err != nil {
		markUnusable(conn)
		return nil, fmt.Errorf("yyp.Client.DoRoundtrip Write: %v", err)
	}
	//t.l.Printf("yyp.Client.DoRoundtrip to %q : written %d", t.S2SName, n)

	b = scratch[0:4]
	_, err = io.ReadFull(conn, b)
	if err != nil {
		markUnusable(conn)
		return nil, fmt.Errorf("yyp.Client.DoRoundtrip Read length: %v", err)
	}
	length := binary.LittleEndian.Uint32(b)
	if length < 10 || length > 64*1024 {
		return nil, fmt.Errorf("yyp.Client.DoRoundtrip invalid length: %v", length)
	}
	b = make([]byte, length-4)
	_, err = io.ReadFull(conn, b)
	if err != nil {
		markUnusable(conn)
		return nil, fmt.Errorf("yyp.Client.DoRoundtrip Read remaining: %v", err)
	}

	var yypresp YYPResp
	if err = yyp.Unmarshal(b,&yypresp);err != nil{
		return
	}
	//var d yyp.DecodeState
	//d.init(b)
	//URI := d.ReadU32()
	//d.ReadU16()

	if yypresp.Uri != uint32(rspURI) {
		return nil, fmt.Errorf("yyp.Client.DoRoundtrip got unexpected response: URI=%v", yypresp.Uri)
	}
	rsp = yypresp.Data
	return
}

func (t *Client) getConn() (net.Conn, error) {
	gIndex, gID, sID := t.findServer()
	if gIndex < 0 {
		return nil, fmt.Errorf("yyp.Client.getConn to %q can not find any server", t.S2SName)
	}

	t.addressLock.RLock()
	if p, ok := t.serverToPool[sID]; ok {
		t.addressLock.RUnlock()
		return p.Get()
	}

	// Create a new pool
	if _, ok := t.addressBook[sID]; !ok {
		err := fmt.Errorf("yyp.Client.getConn to %q can not find serverID=%v in address book", t.S2SName, sID)
		t.l.Printf("Error: %v", err)
		t.addressLock.RUnlock()
		return nil, err
	}

	var target *addr
	if a, ok := t.addressBook[sID][gID][1]; ok { // 优先电信
		target = &a
	} else if a, ok := t.addressBook[sID][gID][2]; ok { // 其次联通
		target = &a
	} else { // 再次随便选
		for _, a := range t.addressBook[sID][gID] {
			target = &a
			break
		}
	}
	t.addressLock.RUnlock()

	if target == nil {
		err := fmt.Errorf("yyp.Client.getConn to %q can not find target address: groupIndex=%v groupID=%v serverID=%v", t.S2SName, gIndex, gID, sID)
		t.l.Printf("Error: %v", err)
		return nil, err
	}

	factory := func() (net.Conn, error) {
		return net.Dial(target.network, target.str)
	}
	p, err := pool.NewChannelPool(0, 6000, factory)
	if err != nil {
		return nil, err
	}
	conn, err := p.Get()

	t.addressLock.Lock()
	t.serverToPool[sID] = p
	t.addressLock.Unlock()

	return conn, err
}

func (t *Client) updateAddress() {
	if t.addressBook == nil {
		t.addressBook = make(map[int64]map[int]map[int]addr)
	}
	for {
		v := <-t.addressNotifier
		leng := int(v.Size())
		if leng == 0 {
			continue
		}
		t.addressLock.Lock()
		for i := 0; i < leng; i++ {
			m := v.Get(i)
			//name := m.GetName()
			server := m.GetServerId()
			group := m.GetGroupId()
			status := m.GetStatus()
			data := s2s.Decode_msg(m.GetData())
			if data == "" {
				t.l.Printf("Error s2s.Decode_msg(data)")
				continue
			}
			var d interface{}
			err := json.Unmarshal([]byte(data), &d)
			if err != nil {
				t.l.Printf("Error json.Unmarshal(%v): %v", data, err)
				continue
			}
			dd := d.(map[string]interface{})
			port := dd["tcp_port"].(string)
			for _, ddd := range dd["iplist"].([]interface{}) {
				dddd := ddd.(map[string]interface{})
				for k, v := range dddd {
					isp := k
					ip := v.(string)
					a, _ := strconv.Atoi(ip)
					aa := net.IPv4(byte(a), byte(a>>8), byte(a>>16), byte(a>>24))
					//t.l.Printf("Debug s2s.Subscribe(%s) got address: %v %v %v %v %v %v %v",	t.S2SName, name, server, group, isp, aa, port, status)
					ispInt, _ := strconv.Atoi(isp)
					if status == 0 {
						if t.addressBook[server] == nil {
							t.addressBook[server] = make(map[int]map[int]addr)
						}
						if t.addressBook[server][group] == nil {
							t.addressBook[server][group] = make(map[int]addr)
						}
						t.addressBook[server][group][ispInt] = addr{network: "tcp", str: aa.String() + ":" + port}
					} else {
						delete(t.addressBook, server)
						if p, ok := t.serverToPool[server]; ok {
							p.Close()
							delete(t.serverToPool, server)
						}
					}
				}
			}
		}
		t.l.Printf("Debug %q addressBook total=%v", t.S2SName, len(t.addressBook))
		// update groupToServerList according to the latest addressBook
		t.groupToServerList = make(map[int][]int64) // we start from a brand new map
		for server, v := range t.addressBook {
			for group := range v {
				if contains(t.Groups, group) {
					list := t.groupToServerList[group]
					t.groupToServerList[group] = append(list, server)
				}
			}
		}
		for g, l := range t.groupToServerList {
			t.l.Printf("Debug %q group %v has %v servers", t.S2SName, g, len(l))
		}
		t.addressLock.Unlock()
	}
}

func (t *Client) findServer() (groupIndex int, groupID int, serverID int64) {
	groupIndex = -1
	groupID = -1
	serverID = -1
	t.addressLock.RLock()
	for i, g := range t.Groups {
		if l, ok := t.groupToServerList[g]; ok {
			groupIndex = i
			groupID = g
			serverID = l[rand.Intn(len(l))]
			break
		}
	}
	t.addressLock.RUnlock()
	return
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func markUnusable(conn net.Conn) {
	if pc, ok := conn.(*pool.PoolConn); ok {
		pc.MarkUnusable()
	}
}
