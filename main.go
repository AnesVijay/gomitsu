package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	botToken = getToken()
)

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
		// TODO: сделать что-то в духе switch-case
		// TODO: нужны сценарии
		//?? - регистрация (через моё одобрение по inline-клавиатуре, которая удаляется после принятия или отказа -> сообщение отмечает как лог моё действие)
		//?? - отправка сообщений челикам из БД (общая отправка всем и только одному, когда это ссылка на подключение)

		if update.Message != nil {
			log.Printf("[%s (ID: %d)]: %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
