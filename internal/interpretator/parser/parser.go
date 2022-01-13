package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/fafnir/internal/interpretator/token"
	"github.com/fafnir/models"
)

func isMathExpression(str string) bool {
	for _, r := range str {
		if !(r >= 40 && r <= 57) {
			return false
		}
	}
	return true
}

func isDate(str string) bool {
	switch strings.ToLower(str) {
	case "вчера":
	case "позавчера":
	case "пн", "вт", "ср", "чт", "пт", "сб", "вс":
		return true
	default:
		// TODO: check format: 12.11
		if true {
			return true
		}
		return false
	}
	return false
}

func Parse(tokens []token.Token) (models.Transaction, error) {
	transaction := models.Transaction{}
	var buffer bytes.Buffer
	for index, token := range tokens {
		if index == 0 || index == 1 {
			if isMathExpression(token.Value) {
				expression, err := govaluate.NewEvaluableExpression(token.Value)
				if err != nil {
					return transaction, err
				}
				result, err := expression.Evaluate(nil)
				if err != nil {
					return transaction, err
				}
				resultStr, ok := result.(float64)
				if !ok {
					return transaction, fmt.Errorf("ошибка преобразования результата вычисления стоимости к числу")
				} else {
					transaction.Price = resultStr
				}
			} else {
				transaction.Category = token.Value
			}
		}
		if index == len(tokens) {
			if isDate(token.Value) {

			} else {
				buffer.WriteString(token.Value)
			}
		} else {
			buffer.WriteString(token.Value)
		}
	}
	transaction.Comment = buffer.String()
	return transaction, nil
}
