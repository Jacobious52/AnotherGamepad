package http

import (
	"net"
	"net/http"
)

const DefaultAddr = ":8080"

type Server struct {
	ln net.Listener

	Handler *Handler

	Addr string
}

func NewServer() *Server {
	return &Server{
		Addr: DefaultAddr,
	}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	go func() {
		http.Serve(s.ln, s.Handler)
	}()

	return nil
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}
