package googlesheets

import (
	"context"
	"fmt"

	"github.com/fafnir/models"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func CreateSpreadsheet(srv *sheets.Service) (*sheets.Spreadsheet, error) {
	sheet, err := srv.Spreadsheets.Create(&sheets.Spreadsheet{}).Do()
	err = fillSpreadsheet(srv, sheet)
	return sheet, err
}

func fillSpreadsheet(srv *sheets.Service, sheet *sheets.Spreadsheet) error {
	// TODO: Generate default spreadsheet content
	return nil
}

func GetSheetsService() (*sheets.Service, error) {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile("./credentials.json"))
	return srv, err
}

// TODO: Add record to spreadsheet
func AddRecord(srv *sheets.Service, spreadsheetId string, transaction models.Transaction) error {
	var vr sheets.ValueRange
	myval := []interface{}{
		transaction.Author,
		transaction.Comment,
		transaction.CommitDate,
		transaction.Price,
		transaction.Category,
		transaction.Type,
	}
	vr.Values = append(vr.Values, myval)
	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, "A2", &vr).InsertDataOption("INSERT_ROWS").ValueInputOption("RAW").Do()
	if err != nil {
		fmt.Printf("Unable to retrieve data from sheet. %v", err)
		return err
	}
	return nil
}
