package handler

import (
	"fmt"

	"github.com/tomassirio/bitcoinTelegram/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := commands.GetPrice()
		b.Send(m.Chat, "BTC's Current price is: U$S "+fmt.Sprint(res))
	}

	commandMap["/historic"] = func(m *tb.Message) {
		res, _ := commands.GetHistoric()
		b.Send(m.Chat, "BTC's Price compared to yesterday is: %"+fmt.Sprint(res))
	}

	return commandMap
}
