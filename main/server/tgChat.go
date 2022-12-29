package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func maina() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("5369772786:AAHFWGS48G8Q3F5HeNzrz3V0RqeAG7rP-bs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID
			var ChatID2 int64
			ChatID2 = 1260057096

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Создаем сообщение

			// и отправляем его
			if ChatID != ChatID2 {
				msg := tgbotapi.NewMessage(ChatID, reply)
				bot.Send(msg)
				msg = tgbotapi.NewMessage(ChatID2, reply)
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(ChatID, reply)
				bot.Send(msg)
			}
		}

	}
}
