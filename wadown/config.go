package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() (Config, error) {
	viper.SetConfigName("wadown.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	for k, v := range viper.AllSettings() {
		fmt.Println("Setting:", k, v)
	}

	var result Config
	if err := viper.Unmarshal(&result); err != nil {
		return result, err
	}
	return result, nil
}

type Config struct {
	KnownPhones map[string]string `mapstructure:"known_phones"`
}

func (cfg *Config) GetUserAlias(userId string) string {
	if cfg == nil || cfg.KnownPhones == nil {
		return userId
	}

	alias, found := cfg.KnownPhones[userId]
	if found {
		return alias
	}
	return userId
}
