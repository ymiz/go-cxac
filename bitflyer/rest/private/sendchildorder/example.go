package sendchildorder

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/bitflyer/config"
	"github.com/ymiz/go-cxac/bitflyer/orderType"
	"github.com/ymiz/go-cxac/bitflyer/product"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/common"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/common/authentication"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/sendchildorder/request"
	"github.com/ymiz/go-cxac/bitflyer/side"
	"github.com/ymiz/go-cxac/bitflyer/timeInForce"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	commonClient := common.NewClient(resty.New(), authentication.NewInformation(conf.Bitflyer.AccessKey, conf.Bitflyer.SecretKey))
	client := NewClient(commonClient)
	resp, body, err := client.Do(request.Body{
		ProductCode:    product.BtcJpy.ToString(), // 現物
		ChildOrderType: orderType.Limit.ToString(),
		Side:           side.Buy.ToString(),
		Price:          3700000,
		Size:           0.001,
		MinuteToExpire: 30,
		TimeInForce:    timeInForce.Gtc.ToString(),
	})
	if err != nil {
		log.Println("sendchildorder error", err.Error())
		return
	}

	log.Println("sendchildorder resp", resp)
	log.Println("sendchildorder response body", body)
}
