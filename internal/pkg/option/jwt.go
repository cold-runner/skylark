package option

import "time"

type Jwt struct {
	Realm            string        `json:"realm"       mapstructure:"realm"`
	TokenHeadName    string        `json:"tokenHeadName" mapstructure:"tokenHeadName"`
	Key              string        `json:"key"         mapstructure:"key"`
	IdentityKey      string        `json:"identityKey" mapstructure:"identityKey"`
	Timeout          time.Duration `json:"timeout"     mapstructure:"timeout"`
	SigningAlgorithm string        `json:"signingAlgorithm" mapstructure:"signingAlgorithm"`
	MaxRefresh       time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
}
