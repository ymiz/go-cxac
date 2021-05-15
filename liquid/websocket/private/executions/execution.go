package executions

import "time"

type RawExecution struct {
	Channel string `json:"channel"`
	Data    string `json:"data"`
	Event   string `json:"event"`
}

type Execution struct {
	Channel string
	Data    Data
	Event   string
}

func NewExecution(channel string, data Data, event string) *Execution {
	return &Execution{Channel: channel, Data: data, Event: event}
}

type RawData struct {
	Id            int     `json:"id"`
	Quantity      float64 `json:"quantity"`
	Price         float64 `json:"price"`
	TakerSide     string  `json:"taker_side"`
	MySide        string  `json:"my_side"`
	CreatedAt     int     `json:"created_at"`
	OrderId       int     `json:"order_id"`
	ClientOrderId string  `json:"client_order_id"`
	Target        string  `json:"target"`
	Timestamp     string  `json:"timestamp"`
}

type Data struct {
	Id            int
	Quantity      float64
	Price         float64
	TakerSide     string
	MySide        string
	CreatedAt     time.Time
	OrderId       int
	ClientOrderId string
	Target        string
	Timestamp     string
}
