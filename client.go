package rpc

import (
	"fmt"
	"log"
	"net"
)

type Program struct {
	Number  uint32
	Version uint32
}

type Client struct {
	prog       Program
	conn       net.Conn
	byteStream bool
}

func Dial(network, address string, prog Program) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return newClient(conn, prog), nil
}

func newClient(conn net.Conn, prog Program) *Client {
	client := &Client{
		prog: prog,
		conn: conn,
	}
	if _, ok := conn.(*net.TCPConn); ok {
		client.byteStream = true
	}
	go func(client *Client) {
		buffer := make([]byte, 1500)
		for {
			n, err := client.conn.Read(buffer)
			if err != nil {
				return
			}

			reply, err := parseReply(buffer[0:n], client.byteStream)
			fmt.Printf("%+v %v\n", reply, err)
		}
	}(client)
	return client
}

func (c *Client) Call(id uint32) {
	call, err := call{
		XId:        2,
		RpcVersion: 2,
		Program:    c.prog,
		Process:    id,
		auth:       0,
		byteStream: c.byteStream,
	}.Seralize()

	log.Println(call, err)
	n, err := c.conn.Write(call)
	log.Println("WROTE", n, err)
}
