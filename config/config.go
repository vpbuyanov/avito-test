package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:"server" validate:"required"`
	Postgres Postgres `mapstructure:"postgres" validate:"required"`
}

type (
	Server struct {
		Address string `mapstructure:"address" validate:"required"`
	}

	Postgres struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     string `mapstructure:"port" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
	}
)

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.yml")

	viper.AddConfigPath("./app")
	viper.AddConfigPath(".")

	viper.SetDefault("postgres.host", "127.0.0.1")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.dbname", "postgres")
	viper.SetDefault("postgres.user", "postgres")
	viper.SetDefault("postgres.password", "postgres")

	viper.SetDefault("server.address", "0.0.0.0:9000")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
