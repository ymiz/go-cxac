package priceLadders

import (
	"github.com/ymiz/go-cxac/liquid/currency/pair"
	"github.com/ymiz/go-cxac/liquid/side"
)

func Generate(code pair.Code, side side.Side) string {
	return "price_ladders_cash_" + code.ToLowerString() + "_" + side.ToLowerString()
}
