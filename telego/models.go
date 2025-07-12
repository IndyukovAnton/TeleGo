package telego

type Chat struct {
	ChatID string
}

type User struct {
	ID           int `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Username     string  `json:"username"`
	LanguageCode string  `json:"language_code"`
}

type SendingMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type Message struct {
	MessageID int       `json:"message_id"`
	From      User      `json:"from"`
	Chat      User      `json:"chat"`
	Date      int       `json:"date"`
	Text      string    `json:"text"`
}

type Update struct {
	UpdateID int      `json:"update_id"`
	Message  Message  `json:"message"`
}

type TelegramResponse struct {
	Ok     bool      `json:"ok"`
	Result []Update  `json:"result"`
}

type Condition func(template string) bool;

type Action func(bot Bot, message Message)