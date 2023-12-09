package config

type QiniuCloud struct {
	AccessKey     string `mapstructure:"accessKey"`     // ak
	SecretKey     string `mapstructure:"secretKey"`     // sk
	Bucket        string `mapstructure:"bucket"`        // 存储桶
	Region        string `mapstructure:"region"`        // 地区
	UseHttps      bool   `mapstructure:"useHttps"`      // 是否启用https
	UseCdnDomains bool   `mapstructure:"useCdnDomains"` // 是否启用CDN加速
	Domain        string `mapstructure:"domain"`        // CDN加速域名
}

type TencentCos struct {
	SecretId   string `mapstructure:"secretId"`
	SecretKey  string `mapstructure:"secretKey"`
	Bucket     string `mapstructure:"bucket"`
	Region     string `mapstructure:"region"`
	Domain     string `mapstructure:"domain"`
	Dictionary string `mapstructure:"dictionary"`
}
