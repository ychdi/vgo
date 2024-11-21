package config

// DbConf 数据数据库配置
type DbConf struct {
	Driver   string `yaml:"driver"`
	Hostname string `yaml:"hostname"`
	HostPort string `yaml:"hostPort"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
