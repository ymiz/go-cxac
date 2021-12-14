package parameter

import (
	"github.com/ymiz/go-cxac/liquid/currency"
	"net/url"
)

type Generator struct {
	parameter url.Values
}

func (g *Generator) Currency(code currency.Code) *Generator {
	g.parameter.Set("currency", code.ToString())
	return g
}

// todo: transaction_type
