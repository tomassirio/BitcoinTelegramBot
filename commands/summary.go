package commands

import (
	"fmt"
	"math"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetSummary() (string, string, error) {
	p, err := utils.GetApiCall()
	l := p.Last
	o := p.Open
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		return fmt.Sprintf("%.2f", p.Last), "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), err
	} else {
		return fmt.Sprintf("%.2f", p.Last), "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), err
	}
}
