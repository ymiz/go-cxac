package get

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/currency"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/rest/private/transactions/get/request/parameter"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	resp, body, err := NewClient(
		common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret),
		resty.New(),
	).Do(parameter.NewGenerator().Currency(currency.Jpy))
	if err != nil {
		log.Println("liquid rest private accounts get error", fmt.Sprintf("%#v", err))
		return
	}

	log.Println("resp", resp)
	log.Println("body", body)
}
