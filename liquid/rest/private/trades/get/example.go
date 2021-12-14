package get

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/rest/private/trades/get/request/parameter"
	"log"
	"strconv"
)

func Example() {
	conf := config.GenerateConfigExample()
	generator := parameter.NewGenerator().Status(parameter.Open).Limit(200)
	resp, body, err := NewClient(
		common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret),
		resty.New(),
	).Do(generator)
	if err != nil {
		log.Println("liquid rest private trades et error: ", err, resp)
		return
	}
	log.Println("response body", body)
	log.Println("resp", resp)

	shortPosition := 0.0
	longPosition := 0.0
	pnl := 0.0

	for _, model := range body.Models {
		f, err := strconv.ParseFloat(model.OpenQuantity, 64)
		log.Println(f)
		if err != nil {
			log.Println("parse float Error", err)
			continue
		}
		if model.Side == "short" {
			shortPosition += f * 10000000
		} else if model.Side == "long" {
			longPosition += f * 10000000
		}
		pnl += model.Pnl
	}
	log.Println("shortPosition is", shortPosition/10000000)
	log.Println("longPosition is", longPosition/10000000)
	log.Println("pnl", pnl)
}
