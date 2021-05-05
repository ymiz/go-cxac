package executions

import "github.com/ymiz/go-cxac/liquid/currency/pair"

func Generate(code pair.Code) string {
	return "executions_cash_" + code.ToLowerString()
}
