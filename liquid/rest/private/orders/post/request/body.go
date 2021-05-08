package request

// @see https://developers.liquid.com/#create-an-order

type Body struct {
	OrderType         string  `json:"order_type"`
	TradingType       string  `json:"trading_type,omitempty"`
	MarginType        string  `json:"margin_type,omitempty"`
	ProductId         int     `json:"product_id"`
	Side              string  `json:"side"`
	Quantity          float64 `json:"quantity"`
	Price             float64 `json:"price,omitempty"`
	PriceRange        float64 `json:"price_range,omitempty"`
	TrailingStopType  string  `json:"trailing_stop_type,omitempty"`
	TrailingStopValue string  `json:"trailing_stop_value,omitempty"`
	ClientOrderId     string  `json:"client_order_id,omitempty"`
	TakerProfit       int     `json:"taker_profit,omitempty"`
	StopLoss          int     `json:"stop_loss,omitempty"`
	LeverageLevel     int     `json:"leverage_level,omitempty"`
	FundingCurrency   string  `json:"funding_currency,omitempty"`
	OrderDirection    string  `json:"order_direction,omitempty"`
}
