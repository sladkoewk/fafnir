package main

import (
	"github.com/fafnir/internal/config"
	"github.com/fafnir/internal/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	infoLog, errorLog := log.Init()
	config, err := config.Init()

	if err != nil {
		errorLog.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.Private.TelegramToken)
	if err != nil {
		errorLog.Fatal(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		text := ""
		if update.Message.Text != "" {
			text = update.Message.Text
			infoLog.Printf("User input: %v", update.Message.Text)
		} else {
			text = config.Messages.Errors.InvalidMessage
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			errorLog.Fatal(err)
		}
	}
}
