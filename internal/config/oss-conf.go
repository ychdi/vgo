package config

// OssConf 阿里云对象存储配置
type OssConf struct {
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Endpoint        string `yaml:"endPoint"`
	BucketName      string `yaml:"bucketName"`
}
