package ticker

import (
	"github.com/gorilla/websocket"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/common/websocket/connection"
	"github.com/ymiz/go-cxac/common/websocket/connection/result"
	"github.com/ymiz/go-cxac/gmo/websocket/subscription"
)

type ErrorHandler struct {
	onUnMarshalError func(err error)
}

func NewErrorHandler(onUnMarshalError func(err error)) *ErrorHandler {
	return &ErrorHandler{onUnMarshalError: onUnMarshalError}
}

type Client struct {
	ch                    chan *Ticker
	wsUrl                 string
	subscriptionParameter subscription.Parameter
	errorHandler          *ErrorHandler
}

func NewClient(
	ch chan *Ticker,
	wsUrl string,
	subscriptionParameter subscription.Parameter,
	errorHandler *ErrorHandler,
) *Client {
	return &Client{ch: ch, wsUrl: wsUrl, subscriptionParameter: subscriptionParameter, errorHandler: errorHandler}
}

func (c *Client) Connect() (r *result.Result) {
	connector := connection.NewConnector(c.wsUrl, func(conn *websocket.Conn) error {
		return conn.WriteJSON(c.subscriptionParameter)
	}, func(message []byte, conn *websocket.Conn) *result.Result {
		var ticker *Ticker
		err := json.Unmarshal(message, &ticker)
		if err != nil {
			if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
				c.errorHandler.onUnMarshalError(err)
			}
			return nil
		}
		c.ch <- ticker
		return nil
	})
	return connector.Connect()
}
