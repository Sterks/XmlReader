package readerxml

// Config ...
type Config struct {
	LogLevel         string `toml:"log_level"`
	ConnectionString string `toml:"ConnectionString"`
}

// NewConfig дефолтная конфигурация
func NewConfig() *Config {
	return &Config{}
}
