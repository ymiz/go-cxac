package get

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	resp, body, err := NewClient(common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret), resty.New()).Do()
	if err != nil {
		log.Println("liquid rest private fiat_accounts get error", err)
		return
	}

	log.Println("resp", resp)
	log.Println("body", body)
	log.Println(body.Models[len(body.Models)-1].Balance)
}
