package configuration

//Configuration ...
type Configuration struct {
	LogLevel         string `toml:"log_level"`
	ConnectionString string `toml:"connection_string"`
	FtpConnect       string `toml:"ftp_connect"`
	RootDir          string `toml:"root_directory"`
	DocType          string `toml:"doc_type"`
	FileDir          string `toml:"file_dir"`
}

// NewConfig дефолтная конфигурация
func NewConfig() *Configuration {
	return &Configuration{
		LogLevel:         "debug",
		ConnectionString: "",
		FtpConnect:       "",
		RootDir:          "",
		DocType:          "",
		FileDir:          "",
	}
}
