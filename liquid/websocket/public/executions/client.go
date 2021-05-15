package executions

import (
	"github.com/gorilla/websocket"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/common/websocket/connection"
	"github.com/ymiz/go-cxac/common/websocket/connection/result"
	"github.com/ymiz/go-cxac/liquid/currency/pair"
	"github.com/ymiz/go-cxac/liquid/websocket/public"
	"github.com/ymiz/go-cxac/liquid/websocket/public/subscription"
	"github.com/ymiz/go-cxac/liquid/websocket/public/subscription/data/channel/executions"
	event2 "github.com/ymiz/go-cxac/liquid/websocket/public/subscription/data/event"
	"github.com/ymiz/go-cxac/liquid/websocket/public/subscription/event"
	"time"
)

type ErrorHandler struct {
	onUnMarshalError func(err error)
}

func NewErrorHandler(onUnMarshalError func(err error)) *ErrorHandler {
	return &ErrorHandler{onUnMarshalError: onUnMarshalError}
}

type Client struct {
	ch               chan *Execution
	currencyPairCode pair.Code
	errorHandler     *ErrorHandler
}

func NewClient(ch chan *Execution, currencyPairCode pair.Code, errorHandler *ErrorHandler) *Client {
	return &Client{ch: ch, currencyPairCode: currencyPairCode, errorHandler: errorHandler}
}

func (c *Client) Connect() (r *result.Result) {
	connector := connection.NewConnector(public.EndPoint, func(conn *websocket.Conn) error {
		return conn.WriteJSON(subscription.Parameter{
			Event: event.Subscribe,
			Data: subscription.Data{
				Channel: executions.Generate(c.currencyPairCode),
				Event:   event2.Updated,
			},
		})
	}, func(message []byte, conn *websocket.Conn) *result.Result {
		var r *RawExecution
		err := json.Unmarshal(message, &r)
		if err != nil {
			if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
				c.errorHandler.onUnMarshalError(err)
			}
			return nil
		}
		var d *RawData
		err = json.Unmarshal([]byte(r.Data), &d)
		if err != nil {
			if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
				c.errorHandler.onUnMarshalError(err)
			}
			return nil
		}
		c.ch <- NewExecution(r.Channel, convertData(d), r.Event)
		return nil
	})
	return connector.Connect()
}

func convertData(r *RawData) Data {
	return Data{
		CreatedAt: time.Unix(int64(r.CreatedAt), 0),
		ID:        r.ID,
		Price:     r.Price,
		Quantity:  r.Quantity,
		TakerSide: r.TakerSide,
		Timestamp: r.Timestamp,
	}
}
