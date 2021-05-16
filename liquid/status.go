package liquid

type Status string

const (
	Live            = Status("live")
	Filled          = Status("filled")
	PartiallyFilled = Status("partially_filled")
	Cancelled       = Status("cancelled")
)

func (s Status) ToString() string {
	return string(s)
}
