package ticker

import (
	"github.com/ymiz/go-cxac/bitflyer/product"
	"github.com/ymiz/go-cxac/bitflyer/websocket/jsonrpc/public"
	"github.com/ymiz/go-cxac/bitflyer/websocket/jsonrpc/subscription"
	"github.com/ymiz/go-cxac/bitflyer/websocket/jsonrpc/subscription/channel/ticker"
	"log"
	"os"
	"os/signal"
)

func Example() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := make(chan *Ticker)
	client := NewClient(c, public.EndPoint, subscription.Parameter{
		Version: "2.0",
		Method:  "subscribe",
		Params:  subscription.Params{Channel: ticker.Generate(product.FxBtcJpy)},
	}, NewErrorHandler(func(err error) {
		log.Println("unmarshal error", err.Error())
	}))
	go client.Connect()

	for {
		select {
		case <-interrupt:
			return
		case t := <-c:
			log.Println("example bf ticker client", t)
		}
	}
}
