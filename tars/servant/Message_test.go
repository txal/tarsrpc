package servant

import (
	"testing"
	"time"
)

func TestMessage(t *testing.T) {
	msg := new(Message)
	msg.Init()
	t.Log(msg.BeginTime)
	time.Sleep(time.Second * 1)
	msg.End()
	t.Log(msg.Cost())
	if msg.Cost() < 1000 || msg.Cost() > 1010 {
		t.Error(msg.Cost())
	}

}
