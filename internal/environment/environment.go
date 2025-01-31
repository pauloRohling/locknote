package environment

type Config struct {
	Server struct {
		Port int `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	} `yaml:"server"`
}
