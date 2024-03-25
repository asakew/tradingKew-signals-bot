package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7027781885:AAFY0C68ssto7VGHeNUzBZG80vjxOeHa_QM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if strings.HasPrefix(update.Message.Text, "/price") {
			coin := "bitcoin"
			resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + coin + "&vs_currencies=usd")
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}

			price := gjson.GetBytes(body, coin+".usd").String()

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Current price of "+strings.Title(coin)+": $"+price)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
