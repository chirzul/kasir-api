package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
	DBURL   string
}

func Load() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if envFileExists(".env") {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
	}

	return &Config{
		AppPort: viper.GetString("PORT"),
		DBURL:   viper.GetString("DATABASE_URL"),
	}
}

func envFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
