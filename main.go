package main

import (
	"github.com/fafnir/internal/bot"
	"github.com/fafnir/internal/config"
	"github.com/fafnir/internal/log"
	"github.com/fafnir/internal/parser"
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
	telegramBotApi.Debug = true
	bot := bot.NewBot(telegramBotApi, config.Messages)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := telegramBotApi.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		infoLog.Printf("User input: %v", update.Message.Text)
		switch update.Message.Text {
		case "/start":
			// TODO: Создание google-таблицы для ведения финансов или ссылка на существующую
			msg.Text = "start"
		case "/table":
			// TODO: Отобразить ссылку на google-таблицу
			msg.Text = "table"
		case "/help":
			// TODO: Справка по основным командам
			msg.Text = "help"
		case "/limits":
			// TODO: Установка лимита расходов на категорию в месяц
			msg.Text = "limits"
		case "/remind":
			// TODO: Установка напоминания об отложенном деле
			msg.Text = "remind"
		case "/auto":
			// TODO: Настройка автоматического проведения транзакций по расписанию (подписки)
			msg.Text = "auto"
		case "/cat":
			// TODO: Показывает все активные категории, которые заведены в таблице и ключевые слова (синонимы) к ним
			msg.Text = "cat"
		case "/bills":
			// TODO: Показывает все счета, которые заведены в таблице и ключевые слова (синонимы) к ним
			msg.Text = "bills"
		case "/today", "/month", "/year":
			// TODO: Показывает статистику за выбранный период
			msg.Text = update.Message.Text
		case "Баланс":
			// TODO: Баланс
			msg.Text = "Баланс"
		default:
			msg.Text = parser.Recognition(update.Message.Text, bot)
		}
		if _, err := telegramBotApi.Send(msg); err != nil {
			errorLog.Fatal(err)
		}
	}
}
