package interpretator

import (
	"fmt"

	"github.com/fafnir/internal/bot"
	"github.com/fafnir/internal/interpretator/lexer"
	"github.com/fafnir/internal/interpretator/parser"
	"github.com/fafnir/models"
)

func GetTransaction(bot *bot.Bot, text string) (models.Transaction, error) {
	tokens := lexer.Tokenization(text)
	fmt.Printf("%v\n", tokens)
	transaction, err := parser.Parse(tokens)
	if err != nil {
		return transaction, fmt.Errorf(bot.Messages.Responses.UnknownCommand, err)
	}
	return transaction, nil
}
