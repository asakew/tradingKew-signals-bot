package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("7027781885:AAGpGKP-RhsIbX1meiFE5XtGKpAVfwbeHKc"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
}
