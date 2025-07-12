package handlers

import (
	"strconv"

	"github.com/telego/telego"
)

func StartHandler(bot telego.Bot, message telego.Message) {
	chat_id := strconv.Itoa(message.Chat.ID)
	text := "Привет, я написанный на собственно-разработанной библиотеки для работы с telegramAPI"

	bot.SendMessage(chat_id, text)
}

func EndHandler(bot telego.Bot, message telego.Message) {
	chat_id := strconv.Itoa(message.Chat.ID)
	text := "До скоро встречи!"

	bot.SendMessage(chat_id, text)
}

func DefaultHandler(bot telego.Bot, message telego.Message) {
	chat_id := strconv.Itoa(message.Chat.ID)
	text := "Я не понимаю команды"

	bot.SendMessage(chat_id, text)
}