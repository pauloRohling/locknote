package environment

import (
	"fmt"
	"time"
)

type Config struct {
	Server struct {
		Port                    int           `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
		GracefulShutdownTimeout time.Duration `yaml:"graceful-shutdown-timeout" env:"SERVER_GRACEFUL_SHUTDOWN_TIMEOUT" env-default:"10s"`
	} `yaml:"server"`
	Log struct {
		Level  LogLevel `yaml:"level" env:"LOG_LEVEL" env-default:"INFO"`
		Groups struct {
			Application  LogLevel `yaml:"application" env:"LOG_GROUP_APPLICATION"`
			Persistence  LogLevel `yaml:"persistence" env:"LOG_GROUP_PERSISTENCE"`
			Presentation LogLevel `yaml:"presentation" env:"LOG_GROUP_PRESENTATION"`
			Security     LogLevel `yaml:"security" env:"LOG_GROUP_SECURITY"`
		} `yaml:"groups"`
	} `yaml:"log"`
	Database struct {
		Host     string `yaml:"host" env:"DATABASE_HOST" env-default:"localhost"`
		Port     int    `yaml:"port" env:"DATABASE_PORT" env-default:"5432"`
		User     string `yaml:"user" env:"DATABASE_USER"`
		Password string `yaml:"password" env:"DATABASE_PASSWORD"`
		Name     string `yaml:"name" env:"DATABASE_NAME"`
	} `yaml:"database"`
	Security struct {
		Auth struct {
			Issuer     string        `yaml:"issuer" env:"SECURITY_AUTH_ISSUER"`
			Expiration time.Duration `yaml:"expiration" env:"SECURITY_AUTH_EXPIRATION" env-default:"1h"`
		} `yaml:"auth"`
		Paseto struct {
			SecretKey string `yaml:"secret-key" env:"SECURITY_PASETO_SECRET_KEY"`
			PublicKey string `yaml:"public-key" env:"SECURITY_PASETO_PUBLIC_KEY"`
		} `yaml:"paseto"`
	} `yaml:"security"`
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

	if config.Security.Paseto.SecretKey == "" {
		return config.getFieldError("security.paseto.secret-key")
	}

	if config.Security.Paseto.PublicKey == "" {
		return config.getFieldError("security.paseto.public-key")
	}

	if config.Security.Auth.Issuer == "" {
		return config.getFieldError("security.auth.issuer")
	}

	return nil
}

func (config *Config) getFieldError(field string) error {
	return fmt.Errorf("field '%s' is required but the value is not provided", field)
}
