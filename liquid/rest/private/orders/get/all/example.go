package all

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/liquid"
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/product"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders/get/all/request/parameter"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	generator := parameter.NewParameterGenerator()
	generator.Status(liquid.Live).ProductId(product.BtcJpy)
	//generator.ProductId(product.BtcJpy)
	resp, body, err := NewClient(common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret), resty.New()).Do(generator)
	if err != nil {
		log.Println("liquid rest private orders get all error: ", err, resp)
		return
	}
	log.Println("response body", body)
	log.Println("resp", resp)
}
