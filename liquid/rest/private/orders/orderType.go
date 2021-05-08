package orders

type OrderType string

const (
	Limit           = OrderType("limit")
	Market          = OrderType("market")
	MarketWithRange = OrderType("market_with_range")
	TrailingStop    = OrderType("trailing_stop")
	LimitPostOnly   = OrderType("limit_post_only")
	Stop            = OrderType("stop")
)

func (t OrderType) ToString() string {
	return string(t)
}
