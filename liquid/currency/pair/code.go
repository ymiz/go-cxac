package pair

import "strings"

type Code string

const (
	BtcJpy = Code("BTCJPY")
	XrpJpy = Code("XRPJPY")
	EthJpy = Code("ETHJPY")
)

func (c Code) ToString() string {
	return string(c)
}

func (c Code) ToLowerString() string {
	return strings.ToLower(string(c))
}
