package symbol

type Symbol string

const (
	BtcUsdt = Symbol("btcusdt")
)

func (s Symbol) ToString() string {
	return string(s)
}
