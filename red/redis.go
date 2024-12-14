package red

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/soonio/tunnel/tunnel"
)

type Red struct {
	conf   Conf
	client *redis.Client
}

func New(c Conf) *Red {
	return &Red{conf: c}
}

func (p *Red) Connect(tunnel *tunnel.Tunnel) error {
	var opt = &redis.Options{
		Addr:         p.conf.Addr,
		Username:     p.conf.Username,
		Password:     p.conf.Password,
		DB:           p.conf.Db,
		ReadTimeout:  -2,
		WriteTimeout: -2,
		Dialer:       tunnel.Client().DialContext,
	}
	p.client = redis.NewClient(opt)
	return p.client.Ping(context.Background()).Err()
}

func (p *Red) Use() *redis.Client {
	return p.client
}

func (p *Red) Close() error {
	if p.client != nil {
		return p.client.Close()
	}
	return nil
}
