package parameter

type Status string

const (
	Open   = Status("open")
	Closed = Status("closed")
)

func (s Status) ToString() string {
	return string(s)
}
