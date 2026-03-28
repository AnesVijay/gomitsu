package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	botToken = getToken()
)

// TODO
// 1. переписать бота с tgbotapi на telegram-bot (там удобнее через хендлеры настраивать поведение бота) или создать map[string]func(), для красивой инициализации команд
// 2. db:
// 2.1 бот должен сам себе создать нужные базы (chatID будет уникальным ключом)
// 2.2 бот должен сам задать туда администратора (его chatID берётся из конфига)

// пример обработки сообщения через эту либу: https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/docs/examples/command-handling.md

func main() {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic()
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// TODO: сделать что-то в духе switch-case => map[string]func()
		// TODO: нужны сценарии
		//?? - регистрация (через моё одобрение по inline-клавиатуре, которая удаляется после принятия или отказа -> сообщение отмечает как лог моё действие)
		//?? - отправка сообщений челикам из БД (общая отправка всем или только одному, когда это ссылка на подключение)

		if update.Message != nil {
			log.Printf("[%s (ID: %d)]: %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
