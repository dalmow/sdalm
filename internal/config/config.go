package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string
	DatabaseUrl string
	BasePath    string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")

	conf := &Config{
		AppPort:     viper.GetString("PORT"),
		DatabaseUrl: viper.GetString("DATABASE_URL"),
		BasePath:    viper.GetString("BASE_PATH"),
	}

	if conf.DatabaseUrl == "" {
		log.Fatal("Database url not provided")
	}

	if conf.BasePath == "" {
		log.Fatal("Base path url not provided")
	}

	return conf, nil
}
