package config

import (
	"log"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
	DBURL   string `mapstructure:"DATABASE_URL"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("no .env file found, using env vars")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "mapstructure"
	}); err != nil {
		return nil, err
	}

	return &cfg, nil
}
