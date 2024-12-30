package config

import (
	"github.com/spf13/viper"
	PkgMongodb "go_graphql_gin/pkg/mongodb"
	"log"
	"os"
)

type Config struct {
	MongoDB PkgMongodb.Config
	Auth0   struct {
		Domain   string
		Audience string
	}
	NotificationService struct {
		URL string
	}
}

func LoadConfig() *Config {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &Config{
		MongoDB: PkgMongodb.Config{
			URI: viper.GetString("MONGO_URI"),
		},
		Auth0: struct {
			Domain   string
			Audience string
		}{
			Domain:   os.Getenv("AUTH0_DOMAIN"),
			Audience: os.Getenv("AUTH0_AUDIENCE"),
		},
	}
}
