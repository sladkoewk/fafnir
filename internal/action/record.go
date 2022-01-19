package action

import (
	"fmt"

	"github.com/fafnir/internal/bot"
	"github.com/fafnir/internal/storage/boltdb"
	"github.com/fafnir/models"
	"google.golang.org/api/sheets/v4"
)

// TODO: Add record to spreadsheet
func AddRecord(b *bot.Bot, chatID int64, transaction models.Transaction) error {
	spreadsheetId, err := b.Storage.Get(chatID, boltdb.Sheet)
	if err != nil {
		return fmt.Errorf(b.Messages.Errors.Default, err)
	}
	var vr sheets.ValueRange
	value := []interface{}{
		transaction.Date,
		transaction.Type,
		transaction.Category,
		transaction.Price,
		transaction.Comment,
		transaction.Author,
		transaction.CommitDate,
		transaction.Сurrency,
	}
	vr.Values = append(vr.Values, value)
	_, err = b.SheetsService.Spreadsheets.Values.Append(spreadsheetId, "A2", &vr).InsertDataOption("INSERT_ROWS").ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet. %v", err)
	}
	return nil
}
