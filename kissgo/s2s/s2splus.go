package s2s

// #cgo CPPFLAGS: -D_GLIBCXX_USE_CXX11_ABI=0
// #cgo LDFLAGS: -ls2sclient  -luuid -lpthread -lrt -lcrypto
import "C"
import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"syscall"
)

/*
@ips: ips[0] is CTL ip, ips[1] is CNC ip.
*/
func EncodeMsg(ips []string, port int) string {
	dip := strings.TrimSpace(ips[0])
	wip := ""
	if len(ips) > 1 {
		wip = strings.TrimSpace(ips[1])
	}

	return Encode_msg(dip, wip, uint(port))
}

type S2sClient struct {
	IMetaServer
	groupId int
	fd      int
	future  chan S2sMetaVector
}

func (s *S2sClient) Register(ips []string, port int) error {
	data := EncodeMsg(ips, port)
	return s.RegisterRaw(data)
}

func (s *S2sClient) RegisterRaw(d string) error {
	if s.SetMine(d) != 0 {
		return errors.New("SetMine Error")
	}
	return nil
}

func (s *S2sClient) UnRegister() error {
	if s.DelMine() != 0 {
		return errors.New("DelMine failed")
	}
	return nil
}

func (s *S2sClient) GetFd() int {
	return s.fd
}

func (s *S2sClient) OnNotify() error {
	v := NewS2sMetaVector()
	status := s.PollNotify(v)
	fmt.Printf("[s2s] pollNotify fd(%d) gid(%d) status(%v)\n", s.fd, s.GetGroupId(), status)
	if s.future != nil {
		s.future <- v
		return nil
	}

	return nil
}

func (s *S2sClient) GetGroupId() int {
	return s.groupId
}

type S2sClientMgr struct {
	token  string
	key    string
	client map[int]*S2sClient

	rwm sync.RWMutex
	efd int
}

func NewS2sClientMgr(token, key string) (*S2sClientMgr, error) {
	efd, err := syscall.EpollCreate(10)
	if err != nil {
		return nil, err
	}
	s := &S2sClientMgr{
		token:  token,
		key:    key,
		client: make(map[int]*S2sClient),
		efd:    efd,
	}
	s.Start()
	return s, nil
}

func (m *S2sClientMgr) Subscribe(name string, gid int, typ int) (chan S2sMetaVector, error) {
	future := make(chan S2sMetaVector, 20)
	ptr := NewMetaServer()

	fd := ptr.Initialize(m.token, m.key, ANY_TYPE)

	ev := syscall.EpollEvent{Events: syscall.EPOLLIN, Fd: int32(fd)}
	if err := syscall.EpollCtl(m.efd, syscall.EPOLL_CTL_ADD, fd, &ev); err != nil {
		//TODO: Clean metaserver
		return future, err
	}

	v := NewSubFilterVector()
	filter := NewSubFilter()
	filter.SetInterestedName(name)
	filter.SetInterestedGroup(gid)
	filter.SetS2sType(typ)
	v.Add(filter)
	if ptr.Subscribe(v) != 0 {
		return future, errors.New("Subscribe error")
	}

	c := &S2sClient{IMetaServer: ptr, fd: fd, groupId: -1, future: future}
	m.rwm.Lock()
	m.client[fd] = c
	m.rwm.Unlock()

	return future, nil
}

// not thread safe
func (m *S2sClientMgr) NewClient(gid int) (*S2sClient, error) {
	return m.NewClient2(gid, MUSIC_PROC)
}

func (m *S2sClientMgr) NewClient2(gid int, t MetaType) (*S2sClient, error) {
	ptr := NewMetaServer()
	if ptr.SetGroupId(uint(gid)) != 0 {
		return nil, errors.New("SetGroupId Error")
	}
	fd := ptr.Initialize(m.token, m.key, t)

	ev := syscall.EpollEvent{Events: syscall.EPOLLIN, Fd: int32(fd)}
	if err := syscall.EpollCtl(m.efd, syscall.EPOLL_CTL_ADD, fd, &ev); err != nil {
		//TODO: Clean metaserver
		return nil, err
	}

	c := &S2sClient{IMetaServer: ptr, fd: fd, groupId: gid}
	m.rwm.Lock()
	m.client[fd] = c
	m.rwm.Unlock()

	return c, nil
}

func (m *S2sClientMgr) Start() {
	go func() {
		for {
			m.Wait()
		}
	}()
}

func (m *S2sClientMgr) Wait() error {
	events := make([]syscall.EpollEvent, 16)
	nfds, err := syscall.EpollWait(m.efd, events, -1)
	if err != nil {

		return err
	}
	for i := 0; i < nfds; i++ {
		fd := int(events[i].Fd)
		m.rwm.RLock()
		c, ok := m.client[fd]
		m.rwm.RUnlock()
		if !ok {
			fmt.Printf("[s2s] pollNotify fd(%d) not found, del it from epoll\n", fd)
			ev := syscall.EpollEvent{}
			if err := syscall.EpollCtl(m.efd, syscall.EPOLL_CTL_DEL, fd, &ev); err != nil {
				fmt.Printf("[s2s] del from epoll failed (%d)\n", fd)
			}
			continue
		}
		err := c.OnNotify()
		if err != nil {
			fmt.Printf("[s2s] Error:%v\n", err)
		}

	}
	return nil
}

var mgr *S2sClientMgr

func InitS2s(token, key string) error {
	var err error
	mgr, err = NewS2sClientMgr(token, key)
	return err
}

var ERR_NOTINIT = errors.New("s2s not init")

func NewS2sClient2(gid int, t MetaType) (*S2sClient, error) {
	if mgr == nil {
		return nil, ERR_NOTINIT
	}
	return mgr.NewClient2(gid, t)
}

func NewS2sClient(gid int) (*S2sClient, error) {
	return NewS2sClient2(gid, MUSIC_PROC)
}

func Register(gid int, ips []string, port int) (*S2sClient, error) {
	c, err := NewS2sClient(gid)
	if err != nil {
		return nil, err
	}
	err = c.Register(ips, port)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Subscribe(name string, gid int) (chan S2sMetaVector, error) {
	return mgr.Subscribe(name, gid, int(MUSIC_PROC))
}

func Subscribe2(name string, gid int, typ int) (chan S2sMetaVector, error) {
	return mgr.Subscribe(name, gid, typ)
}
