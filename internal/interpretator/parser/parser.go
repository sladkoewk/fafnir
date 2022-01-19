package parser

import (
	"bytes"
	"fmt"
	"strings"
	"time"

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

func isDate(str string) (time.Time, bool) {
	nowDate := time.Now()
	switch strings.ToLower(str) {
	case "вчера":
		return nowDate.Add(-24 * time.Hour), true
	case "позавчера":
		return nowDate.Add(-48 * time.Hour), true
	case "пн", "вт", "ср", "чт", "пт", "сб", "вс":
		// TODO: Parse weekday
		return nowDate, true
	default:
		if true {
			// TODO: Parse 21.02
			return nowDate, true
		}
		return nowDate, true
	}
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
		} else if index == len(tokens)-1 {
			date, ok := isDate(token.Value)
			if ok {
				transaction.Date = date.Format("2006-01-02")
			} else {
				buffer.WriteString(fmt.Sprintf(" %s", token.Value))
			}
		} else {
			buffer.WriteString(fmt.Sprintf(" %s", token.Value))
		}
	}
	transaction.Comment = strings.TrimSpace(buffer.String())
	transaction.CommitDate = time.Now().Format("2006-01-02")
	return transaction, nil
}
