package mysql

import "fmt"

type Conf struct {
	Net      string `json:"net" yaml:"net"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Name     string `json:"name" yaml:"name"`
	Params   string `json:"params" yaml:"params"`
}

func (c Conf) Network() string {
	return c.Net
}

func (c Conf) String() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s", c.User, c.Password, c.Net, c.Host, c.Port, c.Name, c.Params)
}
