package product

import "strconv"

type Id int

const (
	BtcJpy = Id(5)
)

func (i Id) ToInt() int {
	return int(i)
}

func (i Id) ToString() string {
	return strconv.Itoa(int(i))
}
