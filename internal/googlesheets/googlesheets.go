package googlesheets

import (
	"context"

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
