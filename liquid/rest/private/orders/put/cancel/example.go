package cancel

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/rest/private/orders/put/cancel/request"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	id := conf.Liquid.CancelId
	resp, body, err := NewClient(
		common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret),
		resty.New(),
	).Do(request.Parameter{Id: id})
	if err != nil {
		log.Println("liquid rest private orders put cancel error:", err, resp)
		return
	}
	log.Println("response body", body)
	log.Println("resp", resp)
}
