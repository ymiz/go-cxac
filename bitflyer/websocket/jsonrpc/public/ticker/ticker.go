package ticker

import "time"

type Ticker struct {
	ProductCode     string    `json:"product_code"`
	State           string    `json:"state"`
	Timestamp       time.Time `json:"timestamp"`
	TickID          int       `json:"tick_id"`
	BestBid         float64   `json:"best_bid"`
	BestAsk         float64   `json:"best_ask"`
	BestBidSize     float64   `json:"best_bid_size"`
	BestAskSize     float64   `json:"best_ask_size"`
	TotalBidDepth   float64   `json:"total_bid_depth"`
	TotalAskDepth   float64   `json:"total_ask_depth"`
	MarketBidSize   float64   `json:"market_bid_size"`
	MarketAskSize   float64   `json:"market_ask_size"`
	Ltp             float64   `json:"ltp"`
	Volume          float64   `json:"volume"`
	VolumeByProduct float64   `json:"volume_by_product"`
}
