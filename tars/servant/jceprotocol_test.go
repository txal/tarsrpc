package servant

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"tarsrpc/jce/taf"
	"tarsrpc/jce_parser/gojce"
	"testing"
)

type helloDispatcher struct {
}

func NewHelloDispatcher() Dispatcher {
	return &helloDispatcher{}
}

type HelloServer interface {
	// Sends a greeting
	TestHello(context.Context, *HelloRequest) (*HelloReply, error)
}

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x94, 0xb8, 0x78, 0x3c,
	0x40, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x8d, 0x8b, 0x0b, 0xaa, 0xa6,
	0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x08, 0xc6,
	0x35, 0xb2, 0xe3, 0x62, 0x05, 0xab, 0x13, 0x32, 0xe5, 0xe2, 0x0c, 0x49, 0x2d, 0x2e, 0x81, 0x70,
	0x84, 0xf5, 0x20, 0xd6, 0x22, 0x5b, 0x23, 0x25, 0x88, 0x2a, 0x58, 0x90, 0x53, 0xa9, 0xc4, 0x90,
	0xc4, 0x06, 0x76, 0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x96, 0x13, 0x77, 0x11, 0xa8, 0x00,
	0x00, 0x00,
}

type HelloServerImpl struct {
}

func (hs *HelloServerImpl) TestHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	panic("handler panic")
}

func (_obj *helloDispatcher) Dispatch(ctx context.Context, _val interface{}, req *taf.RequestPacket) (*taf.ResponsePacket, error) {
	var pbbuf []byte
	_imp := _val.(HelloServer)
	switch req.SFuncName {
	case "TestHello":
		var req_ HelloRequest
		if err := proto.Unmarshal(req.SBuffer, &req_); err != nil {
			return nil, err
		}

		_ret, err := _imp.TestHello(ctx, &req_)
		if err != nil {
			return nil, err
		}

		if pbbuf, err = proto.Marshal(_ret); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("unknow func")
	}
	return &taf.ResponsePacket{
		IVersion:   1,
		IRequestId: req.IRequestId,
		SBuffer:    pbbuf,
		Context:    req.Context,
	}, nil
}

func TestDoDispatch(t *testing.T) {
	raw := HelloRequest{
		Name: "test",
	}
	rawbuf, err := proto.Marshal(&raw)
	if err != nil {
		t.Error("proto.Marshal", err)
	}
	req := &taf.RequestPacket{
		SFuncName: "TestHello",
		SBuffer:   rawbuf,
	}

	os := gojce.NewOutputStream()
	req.WriteTo(os)
	bs := os.ToBytes()
	sbuf := bytes.NewBuffer(nil)
	sbuf.Write(make([]byte, 4))
	sbuf.Write(bs)
	len := sbuf.Len()
	binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(len))

	jceprotocol := NewJceProtocol(NewHelloDispatcher(), &HelloServerImpl{})
	resp, err := jceprotocol.Invoke(context.Background(), sbuf.Bytes())
	if err != nil {
		t.Logf("resp:%v,err:%v", resp, err)
	}
	t.Logf("resp:%v", resp)

	var rspPackage taf.ResponsePacket
	is := gojce.NewInputStream(resp)
	if err = rspPackage.ReadFrom(is); err != nil {
		//this will close the connection
		t.Error("Invoke ReadFrom reqPackage failed", err)
	}
	t.Logf("%v", rspPackage)
}
