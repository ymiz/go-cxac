package product

type Id int

const (
	BtcJpy = Id(5)
)

func (i Id) ToInt() int {
	return int(i)
}
