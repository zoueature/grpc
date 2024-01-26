package grpc

import (
	"github.com/zoueature/config"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Addr string
}

func NewServer(conf config.AppConfig) *Server {
	return &Server{
		Addr: conf.Listen,
	}
}

type RegisterSvc func(s *grpc.Server)

func (serv *Server) Serve(register RegisterSvc) {
	lis, err := net.Listen("tcp", serv.Addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	register(s)
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
