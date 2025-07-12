package main

import (
	"fmt"
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

	message := map[string]string{
		"chat_id": chatID,
		"text": "Hello brother, i`m Working!",
	}

	_ = message

	
	// bot.SendMessage(chatID, message)

	bot.ListenerMessages()
}

func LogLastMessage(message telego.Message) {
	fmt.Printf("User: %s %s (%v)\n", message.From.FirstName, message.From.LastName, message.From.Username)
	fmt.Printf("Message text: %s\n", message.Text)
	fmt.Printf("Date: %v\n", message.Date)
}