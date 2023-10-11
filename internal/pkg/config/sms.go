package config

type TencentSms struct {
	SecretId      string `mapstructure:"secretId"`
	SecretKey     string `mapstructure:"secretKey"`
	Region        string `mapstructure:"region"`
	ApplicationId string `mapstructure:"applicationId"`
	SignName      string `mapstructure:"signName"`
	TemplateId    string `mapstructure:"templateId"`
}
