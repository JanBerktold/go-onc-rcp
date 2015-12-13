package rpc

import (
	"log"
	"net"
)

type Server struct {
}

func (s *Server) addConn(con net.Conn) {
	go func() {
		for {
			buffer := make([]byte, 2560)
			n, err := con.Read(buffer)
			log.Println(n, err)
			if err != nil {
				return
			}
		}
	}()
}

func (s *Server) ListenTCP(network string, laddr *net.TCPAddr) error {
	listener, err := net.ListenTCP(network, laddr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			return err
		}

		s.addConn(conn)
	}

}

func (s *Server) ListenUDP(network string, laddr *net.UDPAddr) error {
	return nil
}
