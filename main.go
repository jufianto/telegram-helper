package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("541876379:AAEeeMMHbwPP2ucTfHehqob8LzaGv7WZCiU")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true
	log.Printf("Authorized on Account %s \n", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 500

	// updates, _ := bot.GetUpdatesChan(updateConfig)

	ChatID := 131109047

	msg := tgbotapi.NewMessage(int64(ChatID), "I Am bro")
	bot.Send(msg)

	// for update := range updates {
	// 	if update.Message == nil {
	// 		continue
	// 	}

	// 	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	// 	fmt.Println("Chat ID", update.Message.Chat.ID)

	// 	bot.Send(msg)
	// }
}
