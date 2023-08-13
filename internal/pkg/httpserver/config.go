package httpserver

type HttpConfig struct {
	Port int `yaml:"port" env:"HTTP_PORT,required"`
}
