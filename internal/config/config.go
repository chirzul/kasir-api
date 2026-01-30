package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `mapstructure:"PORT"`
	DBURL   string `mapstructure:"DATABASE_URL"`
}

func Load() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	var cfg Config
	viper.Unmarshal(&cfg)

	return &cfg, nil
}
