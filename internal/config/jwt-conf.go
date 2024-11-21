package config

// JwtConf Jwt配置
type JwtConf struct {
	AdminKey         string `yaml:"adminKey"`
	AdminSingleLogin int    `yaml:"adminSingleLogin"`
	ApiKey           string `yaml:"apikey"`
	ApiSingleLogin   int    `yaml:"apiSingleLogin"`
	AdminTimeout     int64  `yaml:"adminTimeout"`
	ApiTimeout       int64  `yaml:"apiTimeout"`
}
