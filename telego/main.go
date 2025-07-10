package telego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bot struct {
	Token string
	URL string
}

func NewBot(token string) Bot{
	bot := Bot{
		Token: token,
		URL: "https://api.telegram.org/bot" + token,
	}

	return bot
}

func (b Bot) SendMessage(chatId string, message Message) {
	method := "/sendMessage"

	// Убедимся, что chat_id установлен
	message.ChatID = chatId

	// Преобразуем структуру в JSON
	jsonBody, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Error converting message to JSON:", err)
	}

	// Составляем URL
	fullURL := b.URL + method

	// Создаем запрос
	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(jsonBody))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Устанавливаем заголовок
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending message:", err)
	}
	defer resp.Body.Close()

	// Читаем ответ, чтобы понять, в чём ошибка
	// bodyBytes, _ := io.ReadAll(resp.Body)
	// log.Printf("Status code: %d", resp.StatusCode)
	// log.Printf("Response body: %s", bodyBytes)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d\n", resp.StatusCode)
	}
}

func (b Bot) ListenerMessages() {
	fmt.Println(b.Token)
}