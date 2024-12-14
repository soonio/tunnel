package tunnel

type Conf struct {
	User   string `json:"user" yaml:"user"`     // 用户名
	Secret string `json:"secret" yaml:"secret"` // 密码
	Host   string `json:"host" yaml:"host"`     // 主机IP
	Port   int    `json:"port" yaml:"port"`     // 端口
}
