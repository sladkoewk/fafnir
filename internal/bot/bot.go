package bot

import (
	"github.com/fafnir/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	BotAPI   *tgbotapi.BotAPI
	Messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages) *Bot {
	return &Bot{
		BotAPI:   bot,
		Messages: messages,
	}
}
