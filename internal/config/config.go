package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string
	DatabaseUrl string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")

	conf := &Config{
		AppPort:     viper.GetString("PORT"),
		DatabaseUrl: viper.GetString("DATABASE_URL"),
	}

	if conf.DatabaseUrl == "" {
		log.Fatal("Database url not provided")
	}

	return conf, nil
}
