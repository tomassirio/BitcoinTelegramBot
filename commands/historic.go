package commands

import (
	"fmt"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetHistoric() (string, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	o := p.Open
	his := fmt.Sprintf("%.2f", ((l-o)/o)*100)
	return his, err
}
