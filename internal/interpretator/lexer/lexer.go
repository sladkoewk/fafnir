package lexer

import (
	"fmt"
	"strings"

	"github.com/fafnir/internal/interpretator/token"
)

func Tokenization(text string) []token.Token {
	tokens := []token.Token{}
	words := strings.Split(text, " ")

	for index, word := range words {
		fmt.Printf("index: %v, word: %v\n", index, word)
		tokens = append(tokens, token.Token{Value: strings.TrimSpace(word), Position: index})
	}

	return tokens
}
