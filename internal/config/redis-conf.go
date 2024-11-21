package config

// RedisConf Redis配置
type RedisConf struct {
	Hostname string `yaml:"hostname"`
	HostPort string `yaml:"hostPort"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"database"`
}
