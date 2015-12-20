package rpc

import (
	"log"
	"net"
)

// Server represents a socket which can host multipli RPC programs
type Server struct {
}

// Register is used to make a RPC call available
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

// ListenTCP starts the server on the specified TCP address. Method call
// blocks as long as the server is running.
func (s *Server) ListenTCP(network string, laddr *net.TCPAddr) error {
	listener, err := net.ListenTCP(network, laddr)
	if err != nil {
		return err
	}
	return s.handleListener(listener)
}

// ListenUDP starts the server on the specified UDP address. Method call
// blocks as long as the server is running.
func (s *Server) ListenUDP(network string, laddr *net.UDPAddr) error {
	return nil
}
