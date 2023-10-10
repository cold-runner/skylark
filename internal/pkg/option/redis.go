package option

type Redis struct {
	Host                  string `json:"host"                     mapstructure:"host"                     description:"Cache service host address"`
	Port                  int    `json:"port"                     mapstructure:"port"`
	Username              string `json:"username"                 mapstructure:"username"`
	Password              string `json:"password"                 mapstructure:"password"`
	Database              int    `json:"database"                 mapstructure:"database"`
	MasterName            string `json:"master-name"              mapstructure:"master-name"`
	MaxIdle               int    `json:"optimisation-max-idle"    mapstructure:"optimisation-max-idle"`
	MaxActive             int    `json:"optimisation-max-active"  mapstructure:"optimisation-max-active"`
	Timeout               int    `json:"timeout"                  mapstructure:"timeout"`
	EnableCluster         bool   `json:"enable-cluster"           mapstructure:"enable-cluster"`
	UseSSL                bool   `json:"use-ssl"                  mapstructure:"use-ssl"`
	SSLInsecureSkipVerify bool   `json:"ssl-insecure-skip-verify" mapstructure:"ssl-insecure-skip-verify"`
}
