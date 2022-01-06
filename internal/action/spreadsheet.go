package action

import (
	"fmt"

	"github.com/fafnir/internal/bot"
	"github.com/fafnir/internal/storage/boltdb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetSpreadsheet(b *bot.Bot, message *tgbotapi.Message) error {
	spreadsheetId, err := b.Storage.Get(message.Chat.ID, boltdb.Sheet)
	if err != nil {
		return CheckCreateAccount(b, message)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.Messages.Responses.TableExist, spreadsheetId))
	if _, err := b.BotAPI.Send(msg); err != nil {
		return fmt.Errorf(b.Messages.Errors.Default, err)
	}
	return nil
}
