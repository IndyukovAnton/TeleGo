package telego

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Bot struct {
	Token string
	URL string
	Handlers []Handler
}

func NewBot(token string) Bot{
	bot := Bot{
		Token: token,
		URL: "https://api.telegram.org/bot" + token,
		Handlers : []Handler{},
	}

	return bot
}

func (b *Bot) SendMessage(chat_id string, text string) {
	method := "/sendMessage"

	smessage := SendingMessage{
		ChatID: chat_id,
		Text: text,
	}

	// Преобразуем структуру в JSON
	jsonBody, err := json.Marshal(smessage)
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

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d\n", resp.StatusCode)
	}
}

func (b *Bot) GetUpdate() TelegramResponse {
	method := "/getUpdates"
	fullURL := b.URL + method

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
			log.Fatal("Error creating request:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	var responceData TelegramResponse

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responceData)

	if err != nil {
			log.Fatal("Error decoding JSON:", err)
	}

	return responceData
}

func (b *Bot) RegisterHandler(handler Handler) {
	b.Handlers = append(b.Handlers, handler)
}

func (b *Bot) ListenerMessages() {

	reactionTimeInSecond := 1
	duration := time.Duration(reactionTimeInSecond)*time.Second

	var lastMessage Message;

	for {
		messages := b.GetUpdate()
		result := messages.Result

		if (len(result) == 0) { 
			return
		}

		currentMessage := result[len(result)-1].Message

		if (lastMessage != currentMessage) {
			lastMessage = currentMessage

			for _, handler := range b.Handlers {
				if (handler.Condition(lastMessage.Text)) {
					handler.Action(*b, lastMessage)
					break
				}
			}
		}

		time.Sleep(duration)
	}
}