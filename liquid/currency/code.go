package currency

type Code string

const (
	Jpy  = "JPY"
	Btc  = "BTC"
	Eth  = "ETH"
	Xrp  = "XRP"
	Bch  = "BCH"
	Qash = "QASH"
)

func (c Code) ToString() string {
	return string(c)
}
