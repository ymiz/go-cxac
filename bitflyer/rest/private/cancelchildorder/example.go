package cancelchildorder

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/bitflyer/config"
	"github.com/ymiz/go-cxac/bitflyer/product"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/cancelchildorder/request"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/common"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/common/authentication"
	"log"
)

func Example() {
	conf := config.GenerateConfigExample()
	commonClient := common.NewClient(resty.New(), authentication.NewInformation(conf.Bitflyer.AccessKey, conf.Bitflyer.SecretKey))
	client := NewClient(commonClient)

	resp, err := client.Do(request.Body{
		ProductCode:            product.BtcJpy.ToString(),
		ChildOrderAcceptanceId: conf.Bitflyer.CancelId,
	})
	if err != nil {
		log.Println("cancelchildorder error", err.Error())
		return
	}

	log.Println("cancelchildorder resp", resp)
}
