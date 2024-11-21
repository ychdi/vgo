package config

// CosConf 腾讯云对象存储配置
type CosConf struct {
	SecretID  string `yaml:"secretID"`
	SecretKey string `yaml:"secretKey"`
	Endpoint  string `yaml:"endpoint"`
}
