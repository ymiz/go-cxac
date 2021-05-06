package channel

type Channel string

const (
	Trades = Channel("trades")
	Ticker = Channel("ticker")
)
