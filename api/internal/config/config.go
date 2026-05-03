package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	ServerPort string `mapstructure:"PORT"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: .env file not found, using system environment variables: %v", err)
	}

	err = viper.Unmarshal(&config)
	return
}
