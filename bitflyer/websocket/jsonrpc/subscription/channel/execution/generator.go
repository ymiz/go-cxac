package execution

import "github.com/ymiz/go-cxac/bitflyer/product"

func Generate(productCode product.Code) string {
	return "lightning_executions_" + productCode.ToString()
}
