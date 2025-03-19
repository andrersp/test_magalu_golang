package config

import "github.com/andrersp/favorites/pkg/config"

var cfg Config

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT" required:"true"`
	Api         Api    `mapstructure:"API" required:"true"`
	DB          DB     `mapstructure:"DB" required:"true"`
	Token       Token  `mapstructure:"TOKEN" required:"true"`
}

type Token struct {
	Secret     string `mapstructure:"SECRET" required:"true"`
	Expiration int    `mapstructure:"EXPIRATION" required:"true"`
}

type Api struct {
	Port string `mapstructure:"PORT" required:"true"`
}

func GetApiConfig() Api {
	return cfg.Api
}

type DB struct {
	Host     string `mapstructure:"HOST" required:"true"`
	User     string `mapstructure:"USER" required:"true"`
	Password string `mapstructure:"PASSWORD" required:"true"`
	Name     string `mapstructure:"NAME" required:"true"`
	SSLMode  string `mapstructure:"SSL_MODE" required:"true"`
	Port     int    `mapstructure:"PORT" required:"true"`
}

func GetBDConfig() DB {
	return cfg.DB
}

func GetConfig() Config {
	return cfg
}

func LoadConfig() error {
	if err := config.ReadConfigFromEnvFile(); err != nil {
		return err
	}

	if err := config.ReadConfigFromEnv(&cfg); err != nil {
		return err
	}

	return nil
}
