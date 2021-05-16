package parameter

import (
	"github.com/ymiz/go-cxac/liquid"
	"github.com/ymiz/go-cxac/liquid/currency"
	"github.com/ymiz/go-cxac/liquid/product"
	"net/url"
	"strconv"
)

type Generator struct {
	parameter url.Values
}

func NewParameterGenerator() *Generator {
	return &Generator{
		parameter: map[string][]string{},
	}
}

func (g *Generator) FundingCurrency(code currency.Code) *Generator {
	g.parameter.Set("funding_currency", code.ToString())
	return g
}

func (g *Generator) ProductId(id product.Id) *Generator {
	g.parameter.Set("product_id", id.ToString())
	return g
}

func (g *Generator) Status(status liquid.Status) *Generator {
	g.parameter.Set("status", status.ToString())
	return g
}

func (g *Generator) WithDetails(w int) *Generator {
	g.parameter.Set("with_details", strconv.Itoa(w))
	return g
}

func (g *Generator) Limit(limit int) *Generator {
	g.parameter.Set("limit", strconv.Itoa(limit))
	return g
}

func (g *Generator) Page(page int) *Generator {
	g.parameter.Set("page", strconv.Itoa(page))
	return g
}

func (g *Generator) Generate() url.Values {
	return g.parameter
}
