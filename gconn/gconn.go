package gconn

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RpcConn struct {
	target string
	conn   *grpc.ClientConn
}

func NewRpcConn(target string) *RpcConn {
	return &RpcConn{target: target}
}

func (u *RpcConn) Connect(dialer func(context.Context, string) (net.Conn, error)) error {
	var err error
	u.conn, err = grpc.NewClient(u.target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer))
	return err
}

func (u *RpcConn) Use() *grpc.ClientConn {
	return u.conn
}

func (u *RpcConn) Close() error {
	if u.conn != nil {
		return u.conn.Close()
	}
	return nil
}
