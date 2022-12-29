package main

import (
	"context"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	tgpb "tgbot/pb"
)

var _ tgpb.TGServer = (*GRPCServer)(nil)

type GRPCServer struct{}

func (G GRPCServer) Echo(ctx context.Context, message *tgpb.Message) (*tgpb.Message, error) {
	//подключение к боту по токену
	bot, err := tgbotapi.NewBotAPI("5369772786:AAHFWGS48G8Q3F5HeNzrz3V0RqeAG7rP-bs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	//var ChatID int64
	//var ChatID2 int64
	//ChatID2 = 1023548448
	//ChatID = 1260057096
	var ChatIDArtem int64
	ChatIDArtem = 1789285373

	//чтение обновлений из канала
	log.Printf("[%s] %d %s", message)

	// Ответим пользователю его же сообщением
	reply := message.Message
	// Создаем сообщение
	msg := tgbotapi.NewMessage(ChatIDArtem, reply)
	bot.Send(msg)

	return &tgpb.Message{}, nil
}
