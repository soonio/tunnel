package tunnel

import (
	"context"
	"errors"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

type Tunnel struct {
	conf   Conf
	client *ssh.Client
}

func New(c Conf) *Tunnel {
	return &Tunnel{conf: c}
}

func (s *Tunnel) Connect() error {
	var am ssh.AuthMethod
	if s.conf.Secret != "" {
		am = ssh.Password(s.conf.Secret)
	} else if s.conf.Key != "" {
		signer, err := ssh.ParsePrivateKey([]byte(s.conf.Key))
		if err != nil {
			return err
		}
		am = ssh.PublicKeys(signer)
	} else {
		return errors.New("no auth method")
	}

	config := &ssh.ClientConfig{
		User:            s.conf.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{am},
	}
	var err error
	s.client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port), config)
	return err
}

func (s *Tunnel) Dialer() func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, addr string) (net.Conn, error) {
		return s.client.DialContext(ctx, "tcp", addr)
	}
}

func (s *Tunnel) Client() *ssh.Client {
	return s.client
}

func (s *Tunnel) Close() error {
	if s.client != nil {
		return s.client.Close()
	}
	return nil
}
