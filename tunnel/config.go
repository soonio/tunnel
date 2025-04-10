package tunnel

type Conf struct {
	User   string `json:"user" yaml:"user"`     // 用户名
	Host   string `json:"host" yaml:"host"`     // 主机IP
	Port   int    `json:"port" yaml:"port"`     // 端口
	Secret string `json:"secret" yaml:"secret"` // 密码(二选一)
	Key    string `json:"key" yaml:"key"`       // 私钥(二选一)
}
