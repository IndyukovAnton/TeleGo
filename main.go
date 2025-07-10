package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/telego/telego"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	apiToken := os.Getenv("telegramApiToken")
	chatID := os.Getenv("chatID")

	_ = chatID

	bot := telego.NewBot(apiToken)

	message := telego.Message{
		ChatID: chatID,
		Text: "Hello brother, i`m Working!",
	}
	
	bot.SendMessage(chatID, message)

	// bot.ListenerMessages()
}