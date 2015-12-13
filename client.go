package rpc

import (
	"log"
	"net"
)

type Program struct {
	Number  uint32
	Version uint32
}

type Client struct {
	prog Program
	conn net.Conn
}

func Dial(network, address string, prog Program) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn, prog), nil
}

func NewClient(conn net.Conn, prog Program) *Client {
	client := &Client{
		prog: prog,
		conn: conn,
	}
	go func(client *Client) {
		buffer := make([]byte, 1500)
		for {
			n, err := client.conn.Read(buffer)
			log.Println("READ BUFFER", n, err)
			if err != nil {
				return
			}
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
	}.Seralize()

	log.Println(call, err)
	n, err := c.conn.Write(call)
	log.Println("WROTE", n, err)
}
