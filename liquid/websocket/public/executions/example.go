package executions

import (
	"github.com/ymiz/go-cxac/liquid/currency/pair"
	"log"
	"os"
	"os/signal"
)

func Example() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := make(chan *Execution)
	client := NewClient(c, pair.BtcJpy, nil)
	go client.Connect()

	for {
		select {
		case <-interrupt:
			return
		case execution := <-c:
			log.Println("example executions client", execution)
		}
	}
}
