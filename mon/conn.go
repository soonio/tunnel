package mon

import (
	"net"
	"time"
)

type Conn struct {
	net.Conn
}

// SetWriteDeadline 注释掉原来的设置超时时间bug问题
func (mc *Conn) SetWriteDeadline(_ time.Time) error {
	return nil
}

// SetReadDeadline 注释掉原来的设置超时时间bug问题
func (mc *Conn) SetReadDeadline(_ time.Time) error {
	return nil
}

func (mc *Conn) Read(b []byte) (n int, err error) {
	return mc.Conn.Read(b)
}

func (mc *Conn) Write(b []byte) (n int, err error) {
	return mc.Conn.Write(b)
}
