package trade

import (
	"github.com/gorilla/websocket"
	"github.com/ymiz/go-cxac/binance/symbol"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/common/websocket/connection"
	"github.com/ymiz/go-cxac/common/websocket/connection/result"
)

type ErrorHandler struct {
	onUnMarshalError func(err error)
}

func NewErrorHandler(onUnMarshalError func(err error)) *ErrorHandler {
	return &ErrorHandler{onUnMarshalError: onUnMarshalError}
}

type Client struct {
	ch           chan *Trade
	symbol       symbol.Symbol
	errorHandler *ErrorHandler
}

func NewClient(ch chan *Trade, symbol symbol.Symbol, errorHandler *ErrorHandler) *Client {
	return &Client{ch: ch, symbol: symbol, errorHandler: errorHandler}
}

func (c *Client) Connect() (r *result.Result) {
	connector := connection.NewConnector(c.generateUrl(), func(conn *websocket.Conn) error {
		return nil
	}, func(message []byte, conn *websocket.Conn) *result.Result {
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

func (c *Client) generateUrl() string {
	return "wss://stream.binance.com:9443/ws/" + c.symbol.ToString() + "@trade"
}
