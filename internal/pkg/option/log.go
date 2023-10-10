package option

type Log struct {
	Name              string   `mapstructure:"name"`
	Development       bool     `mapstructure:"development"`
	Level             string   `mapstructure:"level"`
	Format            string   `mapstructure:"format"`
	EnableColor       bool     `mapstructure:"enable-color"`
	DisableCaller     bool     `mapstructure:"disable-caller"`
	DisableStacktrace bool     `mapstructure:"disable-stacktrace"`
	OutputPaths       []string `mapstructure:"output-paths"`
	ErrorOutputPaths  []string `mapstructure:"error-output-paths"`
}
