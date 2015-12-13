package rpc

import (
	"net"
	"testing"
)

func TestStart(t *testing.T) {
	network, err := net.ResolveTCPAddr("tcp", "localhost:2049")
	if err != nil {
		t.Error(err)
	}

	serv := &Server{}
	t.Error(serv.ListenTCP("tcp", network))
}
