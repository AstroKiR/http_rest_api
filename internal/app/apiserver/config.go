package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
}

func GetConfig() *Config {
	return &Config{
		BindAddr: ":80",
	}
}
