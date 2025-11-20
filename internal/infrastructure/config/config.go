package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		DB *DB `yaml:"db"`
	}

	DB struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
)

func New() *Config {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/workspaces/clean-arch-orders/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
