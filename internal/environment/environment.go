package environment

import "fmt"

type Config struct {
	Server struct {
		Port int `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host" env:"DATABASE_HOST" env-default:"localhost"`
		Port     int    `yaml:"port" env:"DATABASE_PORT" env-default:"5432"`
		User     string `yaml:"user" env:"DATABASE_USER"`
		Password string `yaml:"password" env:"DATABASE_PASSWORD"`
		Name     string `yaml:"name" env:"DATABASE_NAME"`
	} `yaml:"database"`
}

func (config *Config) GetDatabaseAddress() string {
	return fmt.Sprintf(
		"%s:%d/%s",
		config.Database.Host, config.Database.Port, config.Database.Name,
	)
}

func (config *Config) GetDatabaseUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User, config.Database.Password,
		config.Database.Host, config.Database.Port,
		config.Database.Name,
	)
}

func (config *Config) validateRequiredFields() error {
	if config.Database.User == "" {
		return config.getFieldError("database.user")
	}

	if config.Database.Password == "" {
		return config.getFieldError("database.password")
	}

	if config.Database.Name == "" {
		return config.getFieldError("database.name")
	}

	return nil
}

func (config *Config) getFieldError(field string) error {
	return fmt.Errorf("field '%s' is required but the value is not provided", field)
}
