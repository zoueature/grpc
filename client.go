package grpc

import (
	"errors"
	"github.com/zoueature/config"
	"google.golang.org/grpc"
)

func NewClient(conf config.RpcConf) (*grpc.ClientConn, error) {
	if conf.Type != config.RpcTypeGrpc {
		return nil, errors.New("Rpc type not match. ")
	}
	conn, err := grpc.Dial(conf.CallHost(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
