package config

var DConfig *DebugConfig

type DebugConfig struct {
	Phone   string // 调试模式下通用手机号，可以自定义
	SmsCode string // 调试模式下通用验证码，可以自定义
}
