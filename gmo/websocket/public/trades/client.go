package trades

import (
	"github.com/gorilla/websocket"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/common/websocket/connection"
	. "github.com/ymiz/go-cxac/common/websocket/connection/result"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription"
)

type ErrorHandler struct {
	onUnMarshalError func(err error)
}

func NewErrorHandler(onUnMarshalError func(err error)) *ErrorHandler {
	return &ErrorHandler{onUnMarshalError: onUnMarshalError}
}

type Client struct {
	ch                    chan *Trade
	wsUrl                 string
	subscriptionParameter subscription.Parameter
	errorHandler          *ErrorHandler
}

func (c *Client) Connect() (r *Result) {
	connector := connection.NewConnector(c.wsUrl, func(conn *websocket.Conn) error {
		return conn.WriteJSON(c.subscriptionParameter)
	}, func(message []byte, conn *websocket.Conn) *Result {
		var trade *Trade
		err := json.Unmarshal(message, &trade)
		if err != nil {
			if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
				c.errorHandler.onUnMarshalError(err)
			}
			return nil
		}
		c.ch <- trade
		return nil
	})
	return connector.Connect()
}

func NewClient(
	ch chan *Trade,
	wsUrl string,
	subscriptionParameter subscription.Parameter,
	errorHandler *ErrorHandler,
) *Client {
	return &Client{ch: ch, wsUrl: wsUrl, subscriptionParameter: subscriptionParameter, errorHandler: errorHandler}
}
