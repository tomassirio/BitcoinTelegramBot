package commands

import "github.com/tomassirio/bitcoinTelegram/utils"

func GetHistoric() (float32, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	bl := p.PriceBeforeLast
	his := (l - bl) / bl
	return his, err
}
