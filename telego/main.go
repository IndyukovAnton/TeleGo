package telego

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func (b *Bot) SendMessage(chatId string, message map[string]string) {
	method := "/sendMessage"

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

			switch strings.ToLower(lastMessage.Text) {
			case "/start":
				message := map[string]string{
					"chat_id": strconv.Itoa(lastMessage.Chat.ID),
					"text": "Привет, я написанный на собственно-разработанной библиотеки для работы с telegramAPI",
				}

				b.SendMessage(strconv.Itoa(lastMessage.Chat.ID), message)
			case "/end":
				message := map[string]string{
					"chat_id": strconv.Itoa(lastMessage.Chat.ID),
					"text": "Пока пока, надеюсь ты ещё зайдёшь!",
				}

				b.SendMessage(strconv.Itoa(lastMessage.Chat.ID), message)
			}
		}

		time.Sleep(duration)
	}
}