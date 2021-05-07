package connection

import (
	"github.com/gorilla/websocket"
	. "github.com/ymiz/go-cxac/common/websocket/connection/result"
	"github.com/ymiz/go-cxac/common/websocket/connection/result/code"
	"os"
	"os/signal"
)

type Connector struct {
	wsUrl                 string
	subscriptionParameter interface{}
	onMessage             func(message []byte)
}

func (c *Connector) Connect() (r *Result) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	r = NewResult(code.Disconnected, nil)

	conn, _, err := websocket.DefaultDialer.Dial(c.wsUrl, nil)
	if err != nil {
		return NewResult(code.DialError, err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			return
		}
		r = NewResult(code.CloseError, err)
	}()

	// NOTE: 少なくともGMOで必要. `サーバーから1分に1回クライアントへpingを送り、3回連続でクライアントからの応答(pong)が無かった場合は、自動的に切断されます。`
	//	 @see https://api.coin.z.com/docs/#public-ws-api
	conn.SetPingHandler(nil)

	err = conn.WriteJSON(c.subscriptionParameter)
	if err != nil {
		return NewResult(code.Disconnected, err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				r = NewResult(code.ReadMessageError, err)
				return
			}
			c.onMessage(message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)
			if err != nil {
				return NewResult(code.SocketCloseError, err)
			}
			select {
			case <-done:
			}
			return NewResult(code.Interrupt, nil)
		}
	}
}

func NewConnector(wsUrl string, subscriptionParameter interface{}, onMessage func(message []byte)) *Connector {
	return &Connector{wsUrl: wsUrl, subscriptionParameter: subscriptionParameter, onMessage: onMessage}
}
