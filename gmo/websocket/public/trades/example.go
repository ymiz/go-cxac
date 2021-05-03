package trades

import (
	"github.com/ymiz/go-cxac/gmo/websocket/public"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription/channel"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription/command"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription/option"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription/symbol"
	"log"
	"os"
	"os/signal"
)

func Example() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := make(chan *Trade)
	client := NewClient(c, public.EndPoint, subscription.Parameter{
		Command: command.Subscribe,
		Channel: channel.Trades,
		Symbol:  symbol.BtcJpy,
		Option:  option.TakerOnly,
	}, nil)
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
