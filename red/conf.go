package red

type Conf struct {
	Addr     string `json:"addr" yaml:"addr"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Db       int    `json:"db" yaml:"db"`
}
