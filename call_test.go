package rpc_test

import (
	"github.com/JanBerktold/go-onc-rpc"
)

var (
	data      = struct{}{}
	bytearray = make([]byte, 0)
)

func ExampleClient_Call() {
	client, _ := rpc.DialTCP("localhost:111", rpc.Program{
		Number:  10000,
		Version: 2,
	})
	// do not forget error checking in the real world

	client.Call(4, rpc.NoReply())
	client.Call(4, rpc.ToStruct(&data))
	client.Call(4, rpc.WithBytes([]byte{0x1, 0x2}), rpc.NoReply())
	client.Call(5, rpc.WithStruct(data), rpc.ToBytes(bytearray))

}
