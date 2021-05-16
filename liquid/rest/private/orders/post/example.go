package post

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/currency"
	"github.com/ymiz/go-cxac/liquid/product"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders/post/request"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders/post/request/leverage"
	"github.com/ymiz/go-cxac/liquid/side"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	requestBody := generateLimitRequestBody()
	//requestBody := generateMarketRequestBody()
	resp, body, err := NewClient(common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret), resty.New()).Do(requestBody)
	if err != nil {
		log.Println("liquid rest private orders post error:", err, resp)
		return
	}
	log.Println("response body", body)
	log.Println("resp", resp)
}

func generateLimitRequestBody() request.Body {
	return leverage.CreateLimitPostOnlyOrderBody(
		product.BtcJpy,
		side.Buy,
		0.0001,
		4500000,
		orders.Twice,
		currency.Jpy,
		orders.NetOut,
	)
}

func generateMarketRequestBody() request.Body {
	return leverage.CreateMarketOrderBody(
		product.BtcJpy,
		side.Buy,
		0.0001,
		orders.Twice,
		currency.Jpy,
		orders.NetOut,
	)
}
