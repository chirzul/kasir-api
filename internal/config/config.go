package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
	DBURL   string `mapstructure:"DATABASE_URL"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Optional defaults
	viper.SetDefault("APP_PORT", "8080")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("no .env file found, using env vars")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
