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
	EnterEmail        string `mapstructure:"enter_email"`
	TableExist        string `mapstructure:"table_exist"`
	UnknownCommand    string `mapstructure:"unknown_command"`
	RecordSaved       string `mapstructure:"record_saved"`
	UnknownCategory   string `mapstructure:"unknown_category"`
	RemovedLastRecord string `mapstructure:"removed_last_record"`
	CreateSpreadsheet string `mapstructure:"create_spreadsheet"`
	CreateTable       string `mapstructure:"create_table"`
}

type Errors struct {
	Default                          string `mapstructure:"default"`
	InvalidEmail                     string `mapstructure:"invalid_email"`
	FailedSaveRecord                 string `mapstructure:"failed_save_record"`
	FailedSaveEmail                  string `mapstructure:"failed_save_email"`
	FailedSaveSpreadsheet            string `mapstructure:"failed_save_spreadsheet"`
	FailedCreateSheet                string `mapstructure:"failed_create_sheet"`
	FailedAuthenticationDriveService string `mapstructure:"failed_authentication_drive_service"`
	FailedShareDocument              string `mapstructure:"failed_share_document"`
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
