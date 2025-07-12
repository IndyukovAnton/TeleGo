package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/telego/handlers"
	"github.com/telego/models"
	"github.com/telego/telego"
)

func NewHandler(condition telego.Condition, action telego.Action) telego.Handler {
	handler := telego.Handler {
		Condition: condition,
		Action: action,
	}

	return handler
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	config := models.Config{
		Token: os.Getenv("telegramApiToken"),
	}

	handlerStart := NewHandler(
		func(template string) bool {
			return strings.ToLower(template) == "/start"
		},
		handlers.StartHandler,
	)

	handlerEnd := NewHandler(
		func(template string) bool {
			return strings.ToLower(template) == "/end"
		},
		handlers.EndHandler,
	)

	defaultHandler := NewHandler(
		func(template string) bool {
			return len(template) > 0
		},
		handlers.DefaultHandler,
	)

	bot := telego.NewBot(config.Token)
	
	bot.RegisterHandler(handlerStart)
	bot.RegisterHandler(handlerEnd)
	bot.RegisterHandler(defaultHandler)

	bot.ListenerMessages()
}