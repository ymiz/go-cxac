package orders

type LeverageLevel int

const (
	Twice = LeverageLevel(2)
)

func (l LeverageLevel) ToInt() int {
	return int(l)
}
