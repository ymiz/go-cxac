package ticker

import (
	"github.com/ymiz/go-cxac/gmo/websocket/public"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription/channel"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription/command"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription/symbol"
	"log"
	"os"
	"os/signal"
)

func Example() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := make(chan *Ticker)
	client := NewClient(c, public.EndPoint, subscription.Parameter{
		Command: command.Subscribe,
		Channel: channel.Ticker,
		Symbol:  symbol.BtcJpy,
	}, nil)
	go client.Connect()

	for {
		select {
		case <-interrupt:
			return
		case ticker := <-c:
			log.Println("example ticker client", ticker)
		}
	}
}
