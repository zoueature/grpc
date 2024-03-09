package grpc

import (
	"context"
	"errors"
	"github.com/zoueature/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func NewClient(conf config.RpcConf) (*grpc.ClientConn, error) {
	if conf.Type != config.RpcTypeGrpc {
		return nil, errors.New("Rpc type not match. ")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() { cancel() }()
	conn, err := grpc.DialContext(
		ctx,
		conf.CallHost(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithIdleTimeout(10*time.Minute),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
