package apiserver

type Config struct{
	BingAddr string `toml:"bing_addr`
	LogLevel string `toml:"log_level`
}

func NewConfig() *Config {
	return &Config{
		BingAddr: ":8080",
		LogLevel: "debug",
	}
}