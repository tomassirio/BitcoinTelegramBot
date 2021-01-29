package commands

import (
	"fmt"
	"math"

	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetHistoric() (string, *tb.Animation, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	o := p.Open
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		g := &tb.Animation{File: tb.FromURL("https://i.pinimg.com/originals/e4/38/99/e4389936b099672128c54d25c4560695.gif")}
		return "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), g, err
	} else {
		g := &tb.Animation{File: tb.FromURL("http://www.brainlesstales.com/bitcoin-assets/images/fan-versions/2015-01-osEroUI.gif")}
		return "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), g, err
	}
}
