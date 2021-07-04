package orderType

type OrderType string

const (
	Limit  = OrderType("LIMIT")
	Market = OrderType("MARKET")
)

func (t OrderType) IsLimit() bool {
	return t == Limit
}

func (t OrderType) IsMarket() bool {
	return t == Market
}

func (t OrderType) ToString() string {
	return string(t)
}
