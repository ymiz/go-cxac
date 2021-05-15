package executions

import "time"

type RawExecution struct {
	Channel string `json:"channel"`
	Data    string `json:"data"` // "data":"{\"created_at\":1620178990,\"id\":424128978,\"price\":5990900.0,\"quantity\":0.002,\"taker_side\":\"buy\",\"timestamp\":\"1620178990.595692\"}"
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
	CreatedAt int     `json:"created_at"`
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
	TakerSide string  `json:"taker_side"`
	Timestamp string  `json:"timestamp"`
}

type Data struct {
	CreatedAt time.Time
	ID        int
	Price     float64
	Quantity  float64
	TakerSide string
	Timestamp time.Time
}
