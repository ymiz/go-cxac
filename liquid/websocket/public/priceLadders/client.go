package priceLadders

import (
	"github.com/gorilla/websocket"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/common/websocket/connection"
	"github.com/ymiz/go-cxac/common/websocket/connection/result"
	"github.com/ymiz/go-cxac/liquid/currency/pair"
	"github.com/ymiz/go-cxac/liquid/side"
	"github.com/ymiz/go-cxac/liquid/websocket/public"
	"github.com/ymiz/go-cxac/liquid/websocket/public/subscription"
	"github.com/ymiz/go-cxac/liquid/websocket/public/subscription/data/channel/priceLadders"
	"github.com/ymiz/go-cxac/liquid/websocket/public/subscription/event"
)

type ErrorHandler struct {
	onUnMarshalError func(err error)
}

func NewErrorHandler(onUnMarshalError func(err error)) *ErrorHandler {
	return &ErrorHandler{onUnMarshalError: onUnMarshalError}
}

type Client struct {
	ch               chan *PriceLadders
	currencyPairCode pair.Code
	side             side.Side
	errorHandler     *ErrorHandler
}

func NewClient(ch chan *PriceLadders, currencyPairCode pair.Code, side side.Side, errorHandler *ErrorHandler) *Client {
	return &Client{ch: ch, currencyPairCode: currencyPairCode, side: side, errorHandler: errorHandler}
}

func (c *Client) Connect() (r *result.Result) {
	connector := connection.NewConnector(public.EndPoint, func(conn *websocket.Conn) error {
		return conn.WriteJSON(subscription.Parameter{
			Event: event.Subscribe,
			Data: subscription.Data{
				Channel: priceLadders.Generate(c.currencyPairCode, c.side),
			},
		})
	}, func(message []byte, conn *websocket.Conn) *result.Result {
		var r *RawPriceLadders
		err := json.Unmarshal(message, &r)
		if err != nil {
			if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
				c.errorHandler.onUnMarshalError(err)
			}
			return nil
		}

		var d *RawData
		err = json.Unmarshal([]byte(`{"raw_price_ladders":`+r.Data+"}"), &d)
		if err != nil {
			if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
				c.errorHandler.onUnMarshalError(err)
			}
			return nil
		}

		c.ch <- NewPriceLadders(r.Channel, d.RawPriceLadders, r.Event)
		return nil
	})
	return connector.Connect()
}
