package response

type Body struct {
	Id                   int64       `json:"id"`
	OrderType            string      `json:"order_type"`
	Quantity             string      `json:"quantity"`
	DiscQuantity         string      `json:"disc_quantity"`
	IcebergTotalQuantity string      `json:"iceberg_total_quantity"`
	Side                 string      `json:"side"`
	FilledQuantity       string      `json:"filled_quantity"`
	Price                float64     `json:"price"`
	CreatedAt            int         `json:"created_at"`
	UpdatedAt            int         `json:"updated_at"`
	Status               string      `json:"status"`
	LeverageLevel        int         `json:"leverage_level"`
	SourceExchange       string      `json:"source_exchange"`
	ProductId            int         `json:"product_id"`
	MarginType           string      `json:"margin_type"`
	TakeProfit           string      `json:"take_profit"`
	StopLoss             string      `json:"stop_loss"`
	TradingType          string      `json:"trading_type"`
	ProductCode          string      `json:"product_code"`
	FundingCurrency      string      `json:"funding_currency"`
	CryptoAccountId      interface{} `json:"crypto_account_id"`
	CurrencyPairCode     string      `json:"currency_pair_code"`
	AveragePrice         float64     `json:"average_price"`
	Target               string      `json:"target"`
	OrderFee             float64     `json:"order_fee"`
	SourceAction         string      `json:"source_action"`
	UnwoundTradeId       interface{} `json:"unwound_trade_id"`
	TradeId              interface{} `json:"trade_id"`
	ClientOrderId        string      `json:"client_order_id"`
}
