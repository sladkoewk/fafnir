package interpretator

import (
	"fmt"

	"github.com/fafnir/internal/bot"
	"github.com/fafnir/internal/interpretator/lexer"
	"github.com/fafnir/internal/interpretator/parser"
)

func GetMessage(bot *bot.Bot, text string) error {
	tokens := lexer.Tokenization(text)
	fmt.Printf("%v\n", tokens)
	transaction, err := parser.Parse(tokens)
	if err != nil {
		return fmt.Errorf(bot.Messages.Responses.UnknownCommand, err)
	}
	fmt.Printf("%+v\n", transaction)
	return fmt.Errorf(bot.Messages.Responses.UnknownCommand)
}
