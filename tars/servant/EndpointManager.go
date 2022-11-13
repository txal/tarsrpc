package servant

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/serialx/hashring"

	"code.com/tars/goframework/jce/servant/taf"
	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
	"code.com/tars/goframework/tars/servant/sd"
	"code.com/tars/goframework/tars/util/endpoint"
)

type EndpointManager struct {
	comm            *Communicator
	objName         string
	refreshInterval int
	directproxy     bool

	mlock           *sync.Mutex
	consistadapters *hashring.HashRing
	adapters        map[string]*AdapterProxy // [ip:port]
	points          map[endpoint.Endpoint]int
	index           []endpoint.Endpoint
	pos             int32
	lastRebootTime  int64
}

func NewEndpointManager(objName string, comm *Communicator) *EndpointManager {
	e := &EndpointManager{
		comm:            comm,
		mlock:           new(sync.Mutex),
		adapters:        make(map[string]*AdapterProxy),
		consistadapters: hashring.NewWithWeights(nil),
		points:          make(map[endpoint.Endpoint]int),
		refreshInterval: comm.Client.refreshEndpointInterval,
	}
	e.setObjName(objName)
	return e
}

func (e *EndpointManager) AddNode(ipPort string) {
	e.consistadapters = e.consistadapters.AddWeightedNode(ipPort, 32)
	appzaplog.Info("AddNode", zap.String("obj", e.objName), zap.String("ipPort", ipPort))
}

func (e *EndpointManager) RemoveNode(ipPort string) {
	e.consistadapters = e.consistadapters.RemoveNode(ipPort)
	appzaplog.Info("RemoveNode", zap.String("obj", e.objName), zap.String("ipPort", ipPort))
}

func (e *EndpointManager) setObjName(objName string) {
	pos := strings.Index(objName, "@")
	if pos > 0 {
		//[direct]
		e.objName = objName[0:pos]
		endpoints := objName[pos+1:]
		e.directproxy = true
		for _, end := range strings.Split(endpoints, ":") {
			e.points[endpoint.Parse(end)] = 0
		}
		e.index = []endpoint.Endpoint{}
		for ep, _ := range e.points {
			e.index = append(e.index, ep)
			e.AddNode(ep.IPPort)
			appzaplog.Debug("consist AddNode", zap.String("ipport", ep.IPPort))
		}
	} else {
		//[proxy] TODO singleton
		appzaplog.Debug("proxy mode", zap.String("objName", objName))
		e.objName = objName
		if objName != "AdminObj" && RedisClient != nil {
			sub := RedisClient.Subscribe(objName)
			appzaplog.Info("sub", zap.String("key", "gracefulReboot"), zap.String("obj", objName))
			go func() {
				for {
					select {
					case msg := <-sub.Channel():
						e.handleRebootPub(msg.Payload)
					}
				}
			}()
		}
		e.findAndSetObj(startFrameWorkComm().sd)
		go func() {
			loop := time.NewTicker(time.Duration(e.refreshInterval) * time.Millisecond)
			for range loop.C {
				//TODO exit
				e.findAndSetObj(startFrameWorkComm().sd)
			}
		}()
	}
}

func (e *EndpointManager) handleRebootPub(msg string) error {
	e.mlock.Lock()
	defer e.mlock.Unlock()
	ipPort, event, err := parsePubMsg(msg)
	if err != nil {
		appzaplog.Error("handleRebootPub parsePubMsg error", zap.Error(err),
			zap.String("key", "gracefulReboot"),
			zap.String("obj", e.objName))
	}
	switch event {
	case "stop":
		now := time.Now().Unix()
		e.lastRebootTime = now
		adp := e.adapters[ipPort]
		if adp != nil {
			e.RemoveNode(ipPort)
			adp.rebootTime = now
			appzaplog.Info("handleRebootPub stop",
				zap.String("key", "gracefulReboot"),
				zap.String("obj", e.objName),
				zap.String("ipPort", ipPort))
		}
	case "start":
		adp := e.adapters[ipPort]
		if adp != nil {
			adp.rebootTime = 0
			e.AddNode(ipPort)
			appzaplog.Info("handleRebootPub start",
				zap.String("key", "gracefulReboot"),
				zap.String("obj", e.objName),
				zap.String("ipPort", ipPort))
		}
	default:
		appzaplog.Error("handleRebootPub unknown event",
			zap.String("key", "gracefulReboot"),
			zap.String("obj", e.objName),
			zap.String("ipPort", ipPort),
			zap.String("event", event))
		return fmt.Errorf("unknown event: %s", event)
	}
	return nil
}

func parsePubMsg(msg string) (string, string, error) {
	parts := strings.Split(msg, "#")
	if len(parts) < 2 {
		appzaplog.Error("parsePubMsg fail", zap.String("key", "gracefulReboot"), zap.String("msg", msg))
		return "", "", fmt.Errorf("unknown pub msg: %s", msg)
	}
	return parts[0], parts[1], nil
}

func (e *EndpointManager) GetNextValidProxy() *AdapterProxy {
	//TODO
	var (
		ep endpoint.Endpoint
	)
	e.mlock.Lock()
	defer e.mlock.Unlock()

	length := len(e.points)
	if length == 0 {
		appzaplog.Error("empty adapter list", zap.Any("endpoint", ep))
		return nil
	}

	for i := 0; i < length; i++ {
		e.pos = (e.pos + 1) % int32(length)
		ep = e.index[e.pos]
		ipport := ep.IPPort
		if adp, ok := e.adapters[ipport]; ok {
			if adp.Available() {
				return adp
			} else {
				continue
			}
		}

		adp, err := e.createProxy(ipport)
		if err != nil {
			appzaplog.Error("create adapter fail", zap.Any("endpoint", ep), zap.Error(err))
			return nil
		}
		return adp
	}

	if length == 1 {
		appzaplog.Warn("single node in circuitbreak open stat", zap.String("ipport", ep.IPPort))
		return nil
	}

	appzaplog.Error("no adapter available", zap.Any("endpoint", ep))
	return nil
}

func (e *EndpointManager) createProxy(ipport string) (*AdapterProxy, error) {
	appzaplog.Debug("create adapter", zap.Any("ipport", ipport))
	host, port, err := net.SplitHostPort(ipport)
	if err != nil {
		return nil, err
	}
	intPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return nil, err
	}
	end := taf.EndpointF{
		Host:  host,
		Port:  int32(intPort),
		Istcp: 1,
	}

	e.adapters[ipport] = NewAdapterProxy(&end, e.comm)
	return e.adapters[ipport], nil
}

func (e *EndpointManager) GetHashProxy(hashcode string) *AdapterProxy {
	e.mlock.Lock()
	length := len(e.points)
	if length == 0 {
		e.mlock.Unlock()
		return nil
	}
	intHashCode, err := strconv.ParseInt(hashcode, 10, 64)
	if err != nil {
		e.mlock.Unlock()
		appzaplog.Error("create adapter fail", zap.String("hashcode", hashcode), zap.Error(err))
		return nil
	}
	pos := intHashCode % int64(length)
	ep := e.index[pos]
	ipport := ep.IPPort
	if adp, ok := e.adapters[ipport]; ok {
		e.mlock.Unlock()
		if adp.Available() {
			return adp
		}
		return nil
	}
	adp, err := e.createProxy(ipport)
	e.mlock.Unlock()
	if err != nil {
		appzaplog.Error("create adapter fail", zap.Any("endpoint", ep), zap.Error(err))
		return nil
	}
	return adp
}

func (e *EndpointManager) GetConsistHashProxy(hashcode string) *AdapterProxy {
	// hot code, dont user defer for lock/unlock
	var (
		localadapter *AdapterProxy
		ipport       string
		ok           bool
		err          error
	)
	e.mlock.Lock()
	defer e.mlock.Unlock()
	node, ok := e.consistadapters.GetNode(hashcode)
	if !ok {
		appzaplog.Warn("GetConsistHashProxy not found", zap.String("hashcode", hashcode))
		return localadapter
	}
	ipport = node

	if localadapter, ok = e.adapters[ipport]; ok {
		if localadapter.Available() {
			return localadapter
		}
		appzaplog.Warn("selected node in circuitbreak open stat,wait 60s", zap.String("ipport", ipport))
		return nil
	}

	localadapter, err = e.createProxy(ipport)
	if err != nil {
		appzaplog.Error("create adapter fail", zap.Any("ipport", ipport), zap.Error(err))
		return nil
	}

	return localadapter
}

func (e *EndpointManager) SelectAdapterProxy(msg IMessage) *AdapterProxy {
	switch {
	case msg.consistHashEnable():
		return e.GetConsistHashProxy(msg.HashCode())
	case msg.hashEnable():
		return e.GetHashProxy(msg.HashCode())
	default:
		return e.GetNextValidProxy()
	}
}

func (e *EndpointManager) findAndSetObj(sdhelper sd.SDHelper) error {
	if sdhelper == nil {
		return NilParamsErr
	}

	activeEp := new([]taf.EndpointF)
	inactiveEp := new([]taf.EndpointF)
	ret, err := sdhelper.FindObjectByIdInSameGroup(e.objName, activeEp, inactiveEp)
	if err != nil {
		appzaplog.Error("find obj end fail", zap.Error(err))
		return err
	}
	appzaplog.Debug("find obj endpoint", zap.String("obj", e.objName), zap.Int32("ret", ret),
		zap.Any("activeEp", *activeEp),
		zap.Any("inactiveEp", *inactiveEp))

	e.mlock.Lock()
	defer e.mlock.Unlock()

	now := time.Now().Unix()
	if now-e.lastRebootTime < int64(e.refreshInterval/1000) {
		appzaplog.Info("find obj skip as rebooting", zap.String("key", "gracefulReboot"), zap.String("obj", e.objName))
		return nil
	}

	if (len(*inactiveEp)) > 0 {
		for _, ep := range *inactiveEp {
			end := endpoint.Taf2endpoint(ep)
			if _, ok := e.points[end]; ok {
				delete(e.points, end)
				ipport := end.IPPort
				delete(e.adapters, ipport)
				e.RemoveNode(ipport)
			}
		}
	}
	if (len(*activeEp)) > 0 {
		tag := int(now)
		e.index = []endpoint.Endpoint{}
		for _, ep := range *activeEp {
			e.points[endpoint.Taf2endpoint(ep)] = tag
			e.index = append(e.index, endpoint.Taf2endpoint(ep))
			e.AddNode(net.JoinHostPort(ep.Host, strconv.FormatInt(int64(ep.Port), 10)))
		}
		for ep, _ := range e.points {
			if e.points[ep] != tag {
				appzaplog.Info("remove ep", zap.Any("endpoint", ep))
				delete(e.points, ep)
				delete(e.adapters, ep.IPPort)
				e.RemoveNode(ep.IPPort)
			}
		}
		for _, adp := range e.adapters {
			adp.rebootTime = 0
		}
	}
	return nil
}

func (e *EndpointManager) GetAvailableProxys() map[string]*AdapterProxy {

	lProxys := map[string]*AdapterProxy{}

	e.mlock.Lock()
	defer e.mlock.Unlock()

	for _, ep := range e.index {
		ipport := ep.IPPort
		if adp, ok := e.adapters[ipport]; ok {
			if adp.Available() {
				lProxys[ipport] = adp
			}
		} else {
			adp, err := e.createProxy(ipport)
			if err != nil {
				appzaplog.Error("GetAvailableProxys adapter fail", zap.Any("endpoint", ep), zap.Error(err))
			} else {
				lProxys[ipport] = adp
			}
		}
	}
	return lProxys
}
