package ticker

import "time"

type Ticker struct {
	Channel   string    `json:"channel"`
	Ask       string    `json:"ask"`
	Bid       string    `json:"bid"`
	High      string    `json:"high"`
	Last      string    `json:"last"`
	Low       string    `json:"low"`
	Symbol    string    `json:"symbol"`
	Timestamp time.Time `json:"timestamp"`
	Volume    string    `json:"volume"`
}
