package configuration

//Configuration ...
type Configuration struct {
	LogLevel         string `toml:"log_level"`
	ConnectionString string `toml:"ConnectionString"`
	Ftp_connect      string `toml:ftp_connect`
}

// NewConfig дефолтная конфигурация
func NewConfig() *Configuration {
	return &Configuration{
		LogLevel:         "debug",
		ConnectionString: "nil",
		Ftp_connect:      "nil",
	}
}
