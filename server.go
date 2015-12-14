package rpc

import (
	"log"
	"net"
)

type Server struct {
}

func (s *Server) Register(prog Program, handler Handler) {

}

func (s *Server) handleConn(conn net.Conn) {
	go func() {
		for {
			buffer := make([]byte, 2560)
			n, err := conn.Read(buffer)
			log.Println(n, err)
			if err != nil {
				return
			}
		}
	}()
}

func (s *Server) handleListener(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			return err
		}

	}
}

func (s *Server) ListenTCP(network string, laddr *net.TCPAddr) error {
	listener, err := net.ListenTCP(network, laddr)
	if err != nil {
		return err
	}
	return s.handleListener(listener)
}

func (s *Server) ListenUDP(network string, laddr *net.UDPAddr) error {
	return nil
}
