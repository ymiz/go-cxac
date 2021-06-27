package trade

import (
	"github.com/ymiz/go-cxac/binance/symbol"
	"log"
	"os"
	"os/signal"
	"time"
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
			ti := time.Unix(0, int64(trade.T1*1000000))
			ti2 := time.Unix(0, int64(trade.E1*1000000))
			log.Println("example trade client", trade, ti, ti2)
		}
	}
}
