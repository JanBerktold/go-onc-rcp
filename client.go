package rpc

import (
	"log"
	"net"
	"sync"
)

type Program struct {
	Number  uint32
	Version uint32
}

type Client struct {
	prog       Program
	conn       net.Conn
	byteStream bool
	curXId     uint32

	repliesMut sync.Mutex
	replies    []chan reply
}

func Dial(network, address string, prog Program) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return newClient(conn, prog), nil
}

func DialTCP(address string, prog Program) (*Client, error) {
	return Dial("tcp", address, prog)
}

func DialUDP(address string, prog Program) (*Client, error) {
	return Dial("udp", address, prog)
}

func newClient(conn net.Conn, prog Program) *Client {
	client := &Client{
		prog:    prog,
		conn:    conn,
		curXId:  1,
		replies: make([]chan reply, 16),
	}
	_, ok := conn.(*net.TCPConn)
	client.byteStream = ok

	go func(client *Client) {
		defer client.conn.Close()
		buffer := make([]byte, 1500)
		for {
			n, err := client.conn.Read(buffer)
			if err != nil {
				return
			}

			reply, err := parseReply(buffer[:n], client.byteStream)
			if err != nil {
				return
			}

			if channel := client.replies[reply.XId]; channel != nil {
				channel <- reply
			} else {
				return
			}
		}
	}(client)
	return client
}

func (c *Client) obtainXId() (uint32, chan reply) {
	for i := 0; i < len(c.replies); i++ {
		if c.replies[i] == nil {
			c.repliesMut.Lock()
			if c.replies[i] == nil {
				channel := make(chan reply, 1)
				c.replies[i] = channel
				c.repliesMut.Unlock()
				return uint32(i), channel
			} else {
				c.repliesMut.Unlock()
				continue
			}
		}
	}

	c.repliesMut.Lock()
	defer c.repliesMut.Unlock()

	// double size of slice
	space := make([]chan reply, len(c.replies)*2)
	copy(space, c.replies)
	c.replies = space
	return c.obtainXId()
}

func (c *Client) Call(proc uint32) ([]byte, error) {
	id, channel := c.obtainXId()
	call, err := call{
		XId:        id,
		RpcVersion: 2,
		Program:    c.prog,
		Process:    proc,
		auth:       0,
		byteStream: c.byteStream,
	}.Seralize()

	if err != nil {
		return nil, err
	}

	if _, err := c.conn.Write(call); err != nil {
		return nil, err
	}

	reply := <-channel
	switch reply.Status.(type) {
	case success:
		return reply.Status.(success).Payload, nil
	default:
		log.Fatal("UNKNOWN SWITCH")
	}
	return nil, nil
}

func (c *Client) CallNoReply(proc uint32) error {
	return nil
}

func (c *Client) CallWithStructReply(proc uint32, reply interface{}) error {
	return nil
}
