package telego

type Message struct {
	ChatID string `json:"chat_id"`
	Text string `json:"text"`
}

type Chat struct {
	ChatID string
}

// type Request struct {

// }