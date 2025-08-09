package config

import (
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Current directory

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &config
}
