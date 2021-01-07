package config

type Config struct {
	host string
	port int
}

func New(config string) Config {
	return Config{"127.0.0.1", 5678}
}
