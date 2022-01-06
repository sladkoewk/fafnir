package main

import (
	"fmt"

	"github.com/fafnir/internal/action"
	"github.com/fafnir/internal/bot"
	"github.com/fafnir/internal/config"
	"github.com/fafnir/internal/googlesheets"
	"github.com/fafnir/internal/log"
	"github.com/fafnir/internal/parser"
	"github.com/fafnir/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	infoLog, errorLog := log.Init()
	config, err := config.Init()
	if err != nil {
		errorLog.Fatal(err)
	}
	telegramBotApi, err := tgbotapi.NewBotAPI(config.Private.TelegramToken)
	if err != nil {
		errorLog.Fatal(err)
	}
	bolt, err := storage.InitBolt()
	if err != nil {
		errorLog.Fatal(err)
	}
	db := storage.NewStorage(bolt)
	sheetsService, err := googlesheets.GetSheetsService()
	if err != nil {
		errorLog.Fatal(err)
	}

	// telegramBotApi.Debug = true

	bot := bot.NewBot(telegramBotApi, config.Messages, db, sheetsService)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := telegramBotApi.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		infoLog.Printf("Ввод пользователя: %v", update.Message.Text)
		err = nil
		err = action.CheckCreateAccount(bot, update.Message)
		if err == nil {
			switch update.Message.Text {
			case "/table":
				err = action.GetSpreadsheet(bot, update.Message)
			case "/help":
				err = action.GetHelp(bot, update.Message)
			case "/limits":
				// TODO: Установка лимита расходов на категорию в месяц
			case "/remind":
				// TODO: Установка напоминания об отложенном деле
			case "/auto":
				// TODO: Настройка автоматического проведения транзакций по расписанию (подписки)
			case "/cat":
				// TODO: Показывает все активные категории, которые заведены в таблице и ключевые слова (синонимы) к ним
			case "/bills":
				// TODO: Показывает все счета, которые заведены в таблице и ключевые слова (синонимы) к ним
			case "/today", "/month", "/year":
				// TODO: Показывает статистику за выбранный период
			case "Баланс":
				// TODO: Баланс
			default:
				err = parser.Recognition(update.Message.Text, bot)
			}
		}
		if err != nil {
			msg.Text = fmt.Sprint(err)
			if _, err := telegramBotApi.Send(msg); err != nil {
				errorLog.Fatal(err)
			}
		}
	}
}
