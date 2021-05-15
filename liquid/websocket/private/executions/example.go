package executions

import (
	"github.com/ymiz/go-cxac/liquid/config"
	"github.com/ymiz/go-cxac/liquid/currency/pair"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"log"
	"os"
	"os/signal"
)

func Example() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	conf := config.GenerateConfigExample()

	c := make(chan *Execution)
	client := NewClient(c, pair.BtcJpy, NewErrorHandler(func(err error) {
		log.Println("onUnMarshalError", err.Error())
	}), common.NewToken(conf.Liquid.ApiTokenId, conf.Liquid.ApiTokenSecret))
	go func() {
		result := client.Connect()
		log.Println("connect result code: ", result.Code(), ", error: ", result.Error())
	}()

	for {
		select {
		case <-interrupt:
			return
		case execution := <-c:
			log.Println("example executions client", execution)
		}
	}
}
