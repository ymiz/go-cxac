package connection

import (
	"github.com/gorilla/websocket"
	. "github.com/ymiz/go-cxac/common/websocket/connection/result"
	"github.com/ymiz/go-cxac/common/websocket/connection/result/code"
	"log"
	"os"
	"os/signal"
)

type Connector struct {
	wsUrl     string
	subscribe func(conn *websocket.Conn) error
	onMessage func(message []byte, conn *websocket.Conn) *Result
}

func NewConnector(
	wsUrl string,
	subscribe func(conn *websocket.Conn) error,
	onMessage func(message []byte, conn *websocket.Conn) *Result,
) *Connector {
	return &Connector{
		wsUrl:     wsUrl,
		subscribe: subscribe,
		onMessage: onMessage,
	}
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
		if err == nil {
			return
		}
		r = NewResult(code.CloseError, err)
	}()

	done := make(chan struct{})
	go func() {
		defer func() {
			c.closeConnectionWithLog(conn)
			close(done)
		}()
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				// TODO: error で返せるようにしたい
				log.Println("ReadMessageError", err)
				break
			}
			re := c.onMessage(message, conn)
			if re != nil {
				r = re
				break
			}
		}
	}()

	// NOTE: 少なくともGMOで必要. `サーバーから1分に1回クライアントへpingを送り、3回連続でクライアントからの応答(pong)が無かった場合は、自動的に切断されます。`
	//	 @see https://api.coin.z.com/docs/#public-ws-api
	conn.SetPingHandler(nil)

	if c.subscribe != nil {
		err := c.subscribe(conn)
		if err != nil {
			r = NewResult(code.Disconnected, err)
			c.closeConnectionWithLog(conn)
		}
	}

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			result := c.closeConnection(conn)
			if result != nil {
				return result
			}
			select {
			case <-done:
			}
			return NewResult(code.Interrupt, nil)
		}
	}
}

func (c *Connector) closeConnection(conn *websocket.Conn) *Result {
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
	if err != nil {
		return NewResult(code.SocketCloseError, err)
	}
	return nil
}

func (c Connector) closeConnectionWithLog(conn *websocket.Conn) {
	result := c.closeConnection(conn)
	if result != nil {
		log.Println("close connection result", result)
	}
}
