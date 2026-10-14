package config

import (
	"os"
	"strings"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct{
	Primary Primary `koanf:"primary" validate:"required"`
	ServerConfig ServerConfig `koanf:"server_config" validate:"required"`
	AuthConfig AuthConfig `koanf:"auth_config" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
	Redis RedisConfig `koanf:"redis" validate:"required"`
}

type Primary struct{
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct{
	ReadTimeout int `koanf:"read_timeout" validate:"required"`
	WriteTimeout int `koanf:"write_timeout" validate:"required"`
	IdleTimeout int `koanf:"idle_timeout" validate:"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origins" validate:"required"`
	Port int `koanf:"port" validate:"required"`
}

type AuthConfig struct{
	SecretKey string `koanf:"secret_key" validate:"required"`
}

type DatabaseConfig struct{
	DSN string `koanf:"dsn" validate:"required"`
}

type RedisConfig struct{
	Address string `koanf:"address" validate:"required"`
}


func LoadConfig() (*Config,error){
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	k := koanf.New(".")

	err := k.Load(env.Provider("BOILERPLATE_", "."),func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "BOILERPLATE_"))
	}, nil)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to load env")
		return nil, err
	}

	mainConfig := &Config{}

	err = k.Unmarshal("",mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to unmarshal config")
	}

	return mainConfig, nil
}