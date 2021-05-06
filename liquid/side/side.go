package side

import "strings"

type Side string

const (
	Buy  = Side("Buy")
	Sell = Side("Sell")
)

func (s Side) ToLowerString() string {
	return strings.ToLower(string(s))
}
