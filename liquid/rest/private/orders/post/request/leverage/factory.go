package leverage

import (
	"github.com/ymiz/go-cxac/liquid/currency"
	"github.com/ymiz/go-cxac/liquid/product"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders/post/request"
	"github.com/ymiz/go-cxac/liquid/side"
)

func CreateLimitOrderBody(
	productId product.Id,
	side side.Side,
	quantity float64,
	price float64,
	leverageLevel orders.LeverageLevel,
	fundingCurrency currency.Code,
	orderDirection orders.OrderDirection,
) request.Body {
	return request.Body{
		OrderType:       orders.Limit.ToString(),
		ProductId:       productId.ToInt(),
		Side:            side.ToLowerString(),
		Quantity:        quantity,
		Price:           price,
		LeverageLevel:   leverageLevel.ToInt(),
		FundingCurrency: fundingCurrency.ToString(),
		OrderDirection:  orderDirection.ToString(),
	}
}

func CreateLimitPostOnlyOrderBody(
	productId product.Id,
	side side.Side,
	quantity float64,
	price float64,
	leverageLevel orders.LeverageLevel,
	fundingCurrency currency.Code,
	orderDirection orders.OrderDirection,
) request.Body {
	return request.Body{
		OrderType:       orders.LimitPostOnly.ToString(),
		ProductId:       productId.ToInt(),
		Side:            side.ToLowerString(),
		Quantity:        quantity,
		Price:           price,
		LeverageLevel:   leverageLevel.ToInt(),
		FundingCurrency: fundingCurrency.ToString(),
		OrderDirection:  orderDirection.ToString(),
	}
}
