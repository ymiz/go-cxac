package product

import "strconv"

type Id int

const (
	BtcJpy = Id(5)
	EtcJpy = Id(29)
	XrpJpy = Id(83)
)

func (i Id) ToInt() int {
	return int(i)
}

func (i Id) ToString() string {
	return strconv.Itoa(int(i))
}
