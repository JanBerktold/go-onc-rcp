package main

import (
	"github.com/JanBerktold/go-onc-rpc"
	"log"
)

func main() {
	client, err := rpc.DialTCP("localhost:111", rpc.Program{
		Number:  100000,
		Version: 2,
	})
	log.Println(client, err)
	client.Call(4)
	for {
	}
}
