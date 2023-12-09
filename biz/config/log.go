package config

// Log contains configuration items related to log.
type Log struct {
	Level       string `json:"level" mapstructure:"level"`
	Format      string `json:"format" mapstructure:"format"`
	OutputPaths string `json:"output-paths" mapstructure:"output-paths"`
}
