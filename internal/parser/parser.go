package parser

import (
	"fmt"

	"github.com/fafnir/internal/bot"
)

func Recognition(text string, bot *bot.Bot) error {

	//TODO: parse text
	//TODO: switch action
	return fmt.Errorf(bot.Messages.Responses.UnknownCommand)
}
