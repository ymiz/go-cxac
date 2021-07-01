package ticker

import "github.com/ymiz/go-cxac/bitflyer/product"

func Generate(productCode product.Code) string {
	return "lightning_ticker_" + productCode.ToString()
}
