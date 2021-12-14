package parameter

import (
	"github.com/ymiz/go-cxac/liquid/currency"
	"net/url"
)

type Generator struct {
	parameter url.Values
}

func NewGenerator() *Generator {
	return &Generator{
		parameter: map[string][]string{},
	}
}
func (g *Generator) Currency(code currency.Code) *Generator {
	g.parameter.Set("currency", code.ToString())
	return g
}

func (g *Generator) Generate() url.Values {
	return g.parameter
}

// todo: transaction_type
