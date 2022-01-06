package account

import (
	"fmt"
	"net/mail"

	"github.com/fafnir/internal/bot"
	"github.com/fafnir/internal/googledrive"
	"github.com/fafnir/internal/googlesheets"
	"github.com/fafnir/internal/log"
	"github.com/fafnir/internal/storage/boltdb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func emailValidation(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CheckCreateAccount(b *bot.Bot, message *tgbotapi.Message) error {
	infoLog, _ := log.Init()
	_, err := b.Storage.Get(message.Chat.ID, boltdb.Email)
	if err != nil {
		infoLog.Printf("У '%v' отсутствует сохраненный email.\n", message.Chat.UserName)
		if message.Text == "/start" {
			msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.Responses.EnterEmail)
			if _, err := b.BotAPI.Send(msg); err != nil {
				return fmt.Errorf(b.Messages.Errors.Default, err)
			}
			return nil
		}
		if emailValidation(message.Text) {
			err := b.Storage.Save(int64(message.Chat.ID), message.Text, boltdb.Email)
			if err != nil {
				return fmt.Errorf(b.Messages.Errors.FailedSaveEmail, err)
			} else {
				infoLog.Printf("Сохранение '%v' корректного email: '%v'.\n", message.Chat.UserName, message.Text)
			}
		} else {
			infoLog.Printf("Введенное '%v' значение '%v' не является email.\n", message.Chat.UserName, message.Text)
			return fmt.Errorf(b.Messages.Errors.InvalidEmail)
		}
	}
	spreadsheetId, err := b.Storage.Get(message.Chat.ID, boltdb.Sheet)
	if err != nil {
		infoLog.Printf("У '%v' отсутствует ссылка на созданную гугл таблицу.\n", message.Chat.UserName)
		msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.Responses.CreateSpreadsheet)
		if _, err := b.BotAPI.Send(msg); err != nil {
			return fmt.Errorf(b.Messages.Errors.Default, err)
		}
		sheet, err := googlesheets.CreateSpreadsheet(b.SheetsService)
		if err != nil {
			return fmt.Errorf(b.Messages.Errors.FailedCreateSheet, err)
		} else {
			infoLog.Printf("Успешно создан документ '%v'.\n", message.Chat.UserName)
		}
		driveService, err := googledrive.GetDriveService()
		if err != nil {
			return fmt.Errorf(b.Messages.FailedAuthenticationDriveService, err)
		} else {
			infoLog.Printf("Успешная авторизация в гугл-драйв сервисе '%v'.\n", message.Chat.UserName)
		}
		err = b.Storage.Save(int64(message.Chat.ID), sheet.SpreadsheetId, boltdb.Sheet)
		if err != nil {
			return fmt.Errorf(b.Messages.FailedSaveSpreadsheet, err)
		} else {
			infoLog.Printf("Успешно сохранен идентификатор документа '%v'.\n", message.Chat.UserName)
		}
		err = googledrive.ShareFile(driveService, sheet.SpreadsheetId, message.Text)
		if err != nil {
			return fmt.Errorf(b.Messages.FailedShareDocument, err)
		} else {
			infoLog.Printf("Документ успешно расшарен по указанному email '%v'.\n", message.Chat.UserName)
		}
		msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.Messages.Responses.CreateTable, sheet.SpreadsheetId))
		if _, err := b.BotAPI.Send(msg); err != nil {
			return fmt.Errorf(b.Messages.Errors.Default, err)
		}
	} else if message.Text == "/start" {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.Messages.Responses.TableExist, spreadsheetId))
		if _, err := b.BotAPI.Send(msg); err != nil {
			return fmt.Errorf(b.Messages.Errors.Default, err)
		} else {
			infoLog.Printf("Проверка наличия email и таблицы '%v' пройдена успешно.\n", message.Chat.UserName)
		}
	}
	return nil
}
