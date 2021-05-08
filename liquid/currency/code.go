package currency

type Code string

const (
	Jpy = "JPY"
)

func (c Code) ToString() string {
	return string(c)
}
