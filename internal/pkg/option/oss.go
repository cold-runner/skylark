package option

type QiniuCloud struct {
	AccessKey     string `mapstructure:"AccessKey"`     // ak
	SecretKey     string `mapstructure:"SecretKey"`     // sk
	Bucket        string `mapstructure:"Bucket"`        // 存储桶
	Region        string `mapstructure:"Region"`        // 地区
	UseHttps      bool   `mapstructure:"UseHttps"`      // 是否启用https
	UseCdnDomains bool   `mapstructure:"UseCdnDomains"` // 是否启用CDN加速
	Domain        string `mapstructure:"Domain"`        // CDN加速域名
}

type TencentCos struct {
	SecretId   string `mapstructure:"secretId"`
	SecretKey  string `mapstructure:"secretKey"`
	Bucket     string `mapstructure:"bucket"`
	Region     string `mapstructure:"region"`
	Domain     string `mapstructure:"domain"`
	Dictionary string `mapstructure:"dictionary"`
}
