package TeleGo

import (
	"fmt"

	"github.com/IndyukovAnton/TeleGo/telego"
)

func LogLastMessage(message telego.Message) {
	fmt.Printf("User: %s %s (%v)\n", message.From.FirstName, message.From.LastName, message.From.Username)
	fmt.Printf("Message text: %s\n", message.Text)
	fmt.Printf("Date: %v\n", message.Date)
}