package side

type Side string

const (
	Buy  = Side("BUY")
	Sell = Side("SELL")
)

func (s Side) IsBuy() bool {
	return s == Buy
}

func (s Side) IsSell() bool {
	return s == Sell
}

func (s Side) ToString() string {
	return string(s)
}
