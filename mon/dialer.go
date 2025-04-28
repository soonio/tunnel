package mon

import (
	"context"
	"golang.org/x/crypto/ssh"
	"net"
)

type Dialer struct {
	ssh *ssh.Client
}

func NewDialer(ssh *ssh.Client) *Dialer {
	return &Dialer{ssh: ssh}
}

func (md *Dialer) DialContext(ctx context.Context, net, addr string) (net.Conn, error) {
	conn, err := md.ssh.DialContext(ctx, net, addr)
	if err != nil {
		return nil, err
	}
	return &Conn{Conn: conn}, nil
}
