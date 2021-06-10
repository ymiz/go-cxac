package trade

import (
	"github.com/ymiz/go-cxac/binance/symbol"
	"log"
	"os"
	"os/signal"
)

func Example() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := make(chan *Trade)
	client := NewClient(c, symbol.BtcUsdt, nil)
	go client.Connect()

	for {
		select {
		case <-interrupt:
			return
		case trade := <-c:
			log.Println("example trade client", trade)
		}
	}
}
