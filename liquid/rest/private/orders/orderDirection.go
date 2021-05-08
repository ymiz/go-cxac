package orders

type OrderDirection string

const (
	NetOut = "netout"
	One    = "one_direction"
	Two    = "two_direction"
)

func (d OrderDirection) ToString() string {
	return string(d)
}
