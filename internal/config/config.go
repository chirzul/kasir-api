package config

import (
	"fmt"
	"os"
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
	fmt.Println("PORT: ", os.Getenv("PORT"))
	fmt.Println("DBURL: ", os.Getenv("DATABASE_URL"))

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
