package trades

import "time"

type Trade struct {
	Channel   string    `json:"channel"`
	Price     string    `json:"price"`
	Side      string    `json:"side"`
	Size      string    `json:"size"`
	Timestamp time.Time `json:"timestamp"`
	Symbol    string    `json:"symbol"`
}
