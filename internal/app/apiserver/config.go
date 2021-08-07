package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func GetConfig() *Config {
	return &Config{
		BindAddr: ":80",
		LogLevel: "error",
	}
}
