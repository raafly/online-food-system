package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const configFilePath = ".env"

func NewAppConfig() (*AppConfig, error) {
	v, err := NewViper(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed initialized viper: %s", err)
	}

	return &AppConfig{
		Fiber: Fiber{
			Port: v.GetString("FIBER_PORT"),
		},
		Postgres: Postgres{
			Host: v.GetString("DATABASE_HOST"),
			Port: v.GetString("DATABASE_PORT"),
			User: v.GetString("DATABASE_USER"),
			Pass: v.GetString("DATABASE_PASS"),
			Name: v.GetString("DATABASE_NAME"),
		},
	}, nil 
}

func NewViper(configFilePath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(configFilePath)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	
	return v, nil
}