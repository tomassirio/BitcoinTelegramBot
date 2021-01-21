package model

type Price struct {
	Last            float32 `json:"last"`
	PriceBeforeLast float32 `json:"price_before_last"`
	Open            float32 `json:"open"`
	High            float32 `json:"high"`
	Low             float32 `json:"low"`
	Vwap            float32 `json:"vwap"`
	Volume          float32 `json:"volume"`
	Bid             float32 `json:"bid"`
	Ask             float32 `json:"ask"`
}

// https://bitex.la/developers#api_ticker
/*
{
  "last":               1230.0,  // Last transaction price
  "price_before_last":  1220.0,  // Helps you tell if price is going up or down.
  "open":        1198.45875559,  // What the price was 24 hours ago.
  "high":               1230.0,  // Highest price for the past 24 hours.
  "low":          1193.2507548,  // Lowest price for the past 24 hours.
  "vwap":        1208.57944642,  // Volume-Weighted Average Price for the past 24 hours.
  "volume":        16.45315074,  // Transacted volume for the last 24 hours.
  "bid":          1226.5583985,  // Highest current buy order.
  "ask":         1235.71481927   // Lowest current ask order.
}
*/
