package product

type Code string

const (
	FxBtcJpy = Code("FX_BTC_JPY")
	BtcJpy   = Code("BTC_JPY")
)

func (c Code) IsFxBtcJpy() bool {
	return c == FxBtcJpy
}

func (c Code) IsBtcJpy() bool {
	return c == BtcJpy
}

func (c Code) ToString() string {
	return string(c)
}
