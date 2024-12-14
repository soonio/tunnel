package mysql

import (
	"context"
	"database/sql"
	"net"

	"github.com/go-sql-driver/mysql"
)

type Databases struct {
	db   *sql.DB
	conf Conf
}

func New(c Conf) *Databases {
	return &Databases{conf: c}
}

func (d *Databases) Connect(dialer func(context.Context, string) (net.Conn, error)) error {
	mysql.RegisterDialContext(d.conf.Network(), dialer)

	var err error
	d.db, err = sql.Open("mysql", d.conf.String())
	return err
}

func (d *Databases) Use() *sql.DB {
	return d.db
}

func (d *Databases) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}
