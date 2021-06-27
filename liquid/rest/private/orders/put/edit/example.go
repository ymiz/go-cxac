package edit

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders/put/edit/request"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	id := conf.Liquid.ChangeId
	price := conf.Liquid.ChangePrice
	q := conf.Liquid.ChangeQuantity
	order := request.Order{
		Quantity: q,
		Price:    price,
	}
	body := request.Body{Order: order}
	resp, responseBody, err := NewClient(
		common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret),
		resty.New(),
	).Do(request.Parameter{Id: id}, body)

	if err != nil {
		log.Println("liquid rest private orders put edit error:", err)
		return
	}

	log.Println("response body", responseBody)
	log.Println("resp", resp)
}
