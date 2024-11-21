package config

// QueueConf Queue配置
type QueueConf struct {
	Enable      int    `yaml:"enable"`
	Hostname    string `yaml:"hostname"`
	HostPort    string `yaml:"hostPort"`
	UserName    string `yaml:"username"`
	Password    string `yaml:"password"`
	DB          int    `yaml:"db"`
	Concurrency int    `yaml:"concurrency"`
}
