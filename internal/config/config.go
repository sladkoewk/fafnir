package config

import (
	"github.com/spf13/viper"
)

type Private struct {
	TelegramToken string `mapstructure:"telegram_token"`
}

type Messages struct {
	Responses
	Errors
}

type Config struct {
	Messages Messages
	Private  Private
}

type Responses struct {
	Start string `mapstructure:"start"`
}

type Errors struct {
	Default        string `mapstructure:"default"`
	InvalidMessage string `mapstructure:"invalid_message"`
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("telegram_token", &cfg.Private.TelegramToken); err != nil {
		return err
	}

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
