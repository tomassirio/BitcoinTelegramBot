package commands

import (
	"fmt"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetPrice() (string, error) {
	p, err := utils.GetApiCall()
	return fmt.Sprintf("%.2f", p.Last), err
}
