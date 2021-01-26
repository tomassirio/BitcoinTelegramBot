package commands

import (
	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetPrice() (float32, error) {
	p, err := utils.GetApiCall()
	return p.Last, err
}
