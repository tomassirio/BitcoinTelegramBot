package handler

import (
	"github.com/tomassirio/bitcoinTelegram/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := commands.GetPrice()
		b.Send(m.Chat, "BTC's Current price is: U$S "+res)
	}

	commandMap["/historic"] = func(m *tb.Message) {
		res, g, _ := commands.GetHistoric()
		b.Send(m.Chat, "BTC's Price compared to yesterday is: "+res)
		b.Send(m.Chat, g)
	}

	commandMap["/summary"] = func(m *tb.Message) {
		p, h, _ := commands.GetSummary()
		b.Send(m.Chat, "BTC's Current price is: U$S "+p+"\nBTC's Price compared to yesterday is: "+h)
	}

	return commandMap
}
