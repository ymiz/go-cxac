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
		log.Println("liquid rest private trades et error: ", err, resp)
		return
	}
	log.Println("response body", body)
	log.Println("resp", resp)
}
