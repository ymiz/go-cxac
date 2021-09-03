package parameter

import (
	"github.com/ymiz/go-cxac/liquid/currency"
	"net/url"
	"strconv"
)

type Generator struct {
	parameter url.Values
}

func NewGenerator() *Generator {
	return &Generator{
		parameter: map[string][]string{},
	}
}

func (g *Generator) FundingCurrency(code currency.Code) *Generator {
	g.parameter.Set("funding_currency", code.ToString())
	return g
}

func (g *Generator) Status(status Status) *Generator {
	g.parameter.Set("status", status.ToString())
	return g
}

func (g *Generator) Limit(limit int) *Generator {
	g.parameter.Set("limit", strconv.Itoa(limit))
	return g
}

func (g *Generator) Generate() url.Values {
	return g.parameter
}
