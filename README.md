# TeleGo
Go library for working with Telegram

download:

```
go get github.com/IndyukovAnton/TeleGo@v0.1.1
```


### Example:
```
bot := telego.NewBot("<token>")

handlerStart := TeleGo.NewHandler(
	func(template string) bool {
		return strings.ToLower(template) == "/start"
	},
	handlers.StartHandler,
)

handlerEnd := TeleGo.NewHandler(
	func(template string) bool {
		return strings.ToLower(template) == "/end"
	},
	handlers.EndHandler,
)

bot.RegisterHandler(handlerStart)
bot.RegisterHandler(handlerEnd)

bot.ListenerMessages()
```



# Bot

Use telego.NewBot for getting object Bot

```

func NewBot(token string) Bot{
	bot := Bot{
		Token: token,
		URL: "https://api.telegram.org/bot" + token,
		Handlers : []Handler{},
	}

	return bot
}

```

```
type Bot struct {
	Token string
	URL string
	Handlers []Handler
}
```

## Methods:


```
SendMessage(chat_id string, text string) - Sending message
```

```
GetUpdate() TelegramResponse - Returning latest messages
```

```
RegisterHandler(handler Handler) - Registering new handler
```

```
ListenerMessages() - Event loop for listening and reaction on new messages, using registering handlers. One reaction in second
```