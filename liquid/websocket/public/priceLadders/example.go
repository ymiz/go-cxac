package priceLadders

import (
	"github.com/ymiz/go-cxac/liquid/currency/pair"
	"github.com/ymiz/go-cxac/liquid/side"
	"log"
	"os"
	"os/signal"
)

func Example() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := make(chan *PriceLadders)
	client := NewClient(c, pair.BtcJpy, side.Buy, NewErrorHandler(func(err error) {
		log.Println("priceLadders example onUnmarshalError", err.Error())
	}))
	go client.Connect()

	for {
		select {
		case <-interrupt:
			return
		case priceLadders := <-c:
			log.Println("example priceLadders client", priceLadders)
		}
	}
}
