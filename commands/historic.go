package commands

import "github.com/tomassirio/bitcoinTelegram/utils"

func GetHistoric() (float32, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	o := p.Open
	his := (l - o) / o
	return his, err
}
