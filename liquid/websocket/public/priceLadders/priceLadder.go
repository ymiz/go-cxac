package priceLadders

type RawPriceLadders struct {
	Channel string `json:"channel"`
	Data    string `json:"data"`
	Event   string `json:"event"`
}

type RawData struct {
	RawPriceLadders [][]string `json:"raw_price_ladders"`
}

type PriceLadders struct {
	Channel string
	Data    [][]string // [["6253363.00000", "0.00220896"]] のように, [price,size] が入っている.
	Event   string
}

func NewPriceLadders(channel string, data [][]string, event string) *PriceLadders {
	return &PriceLadders{Channel: channel, Data: data, Event: event}
}
