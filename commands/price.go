package commands

import (
	"encoding/json"
	"net/http"

	"github.com/tomassirio/bitcoinTelegram/model"
)

func GetPrice() (float32, error) {
	resp, err := http.Get("https://bitex.la/api-v1/rest/btc_usd/market/ticker")
	p := &model.Price{}
	if err != nil {
		return 0.0, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p.Last, err
}
