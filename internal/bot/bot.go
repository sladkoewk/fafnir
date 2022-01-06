package bot

import (
	"github.com/fafnir/internal/config"
	"github.com/fafnir/internal/storage/boltdb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/sheets/v4"
)

type Bot struct {
	BotAPI        *tgbotapi.BotAPI
	Messages      config.Messages
	Storage       boltdb.AppStorage
	SheetsService *sheets.Service
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages, storage boltdb.AppStorage, sheetsService *sheets.Service) *Bot {
	return &Bot{
		BotAPI:        bot,
		Messages:      messages,
		Storage:       storage,
		SheetsService: sheetsService,
	}
}
