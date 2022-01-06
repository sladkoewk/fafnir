package action

import (
	"fmt"

	"github.com/fafnir/internal/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetHelp(b *bot.Bot, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.Responses.Help)
	msg.ParseMode = tgbotapi.ModeMarkdown
	if _, err := b.BotAPI.Send(msg); err != nil {
		return fmt.Errorf(b.Messages.Errors.Default, err)
	}
	return nil
}
