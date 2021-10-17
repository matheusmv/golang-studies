package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func LoadConfigurations() (Configurations, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var cfg Configurations
	var err error

	err = viper.ReadInConfig()

	if err != nil {
		errorMessage := fmt.Sprintf("Error reading config file, %s", err)
		return Configurations{}, errors.New(errorMessage)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		errorMessage := fmt.Sprintf("Unable to decode into struct, %v", err)
		return Configurations{}, errors.New(errorMessage)
	}

	return cfg, nil
}
