package grpc

import (
	"github.com/zoueature/config"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Addr string
}

func NewServer(conf *config.AppConfig) *Server {
	return &Server{
		Addr: conf.GrpcListen,
	}
}

type RegisterSvc func(s *grpc.Server)

func (serv *Server) Serve(register RegisterSvc) error {
	listener, err := net.Listen("tcp", serv.Addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	register(s)
	err = s.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
