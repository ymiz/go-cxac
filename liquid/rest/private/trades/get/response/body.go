package response

type Body struct {
	Models []struct {
		ID                int         `json:"id"`
		CurrencyPairCode  string      `json:"currency_pair_code"`
		Status            string      `json:"status"`
		Side              string      `json:"side"`
		MarginUsed        string      `json:"margin_used"`
		OpenQuantity      string      `json:"open_quantity"`
		CloseQuantity     string      `json:"close_quantity"`
		Quantity          string      `json:"quantity"`
		ClaimQuantity     string      `json:"claim_quantity"`
		LeverageLevel     int         `json:"leverage_level"`
		ProductCode       string      `json:"product_code"`
		ProductID         int         `json:"product_id"`
		OriginalOpenPrice string      `json:"original_open_price"`
		OpenPrice         string      `json:"open_price"`
		ClosePrice        string      `json:"close_price"`
		TraderID          int         `json:"trader_id"`
		OpenPnl           float64     `json:"open_pnl"`
		ClosePnl          string      `json:"close_pnl"`
		Pnl               float64     `json:"pnl"`
		StopLoss          string      `json:"stop_loss"`
		TakeProfit        string      `json:"take_profit"`
		MarginType        string      `json:"margin_type"`
		FundingCurrency   string      `json:"funding_currency"`
		TotalInterest     string      `json:"total_interest"`
		CreatedAt         int         `json:"created_at"`
		UpdatedAt         int         `json:"updated_at"`
		TotalFee          float64     `json:"total_fee"`
		CloseFee          string      `json:"close_fee"`
		LiquidationPrice  interface{} `json:"liquidation_price"`
		MaintenanceMargin interface{} `json:"maintenance_margin"`
		TradingType       string      `json:"trading_type"`
		LastSettlementAt  interface{} `json:"last_settlement_at"`
		OrderID           int64       `json:"order_id"`
	} `json:"models"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}
