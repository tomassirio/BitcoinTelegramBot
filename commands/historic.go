package commands

import (
	"math"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetHistoric() (float32, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	o := p.Open
	his := float32(math.Round(float64((l-o)/o) * 100))
	return his, err
}
