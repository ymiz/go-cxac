package executions

import (
	"errors"
	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/common/websocket/connection"
	"github.com/ymiz/go-cxac/common/websocket/connection/result"
	"github.com/ymiz/go-cxac/common/websocket/connection/result/code"
	"github.com/ymiz/go-cxac/liquid/currency/pair"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/websocket/private"
	"github.com/ymiz/go-cxac/liquid/websocket/private/auth"
	"github.com/ymiz/go-cxac/liquid/websocket/private/auth/data"
	headers3 "github.com/ymiz/go-cxac/liquid/websocket/private/auth/data/headers"
	event2 "github.com/ymiz/go-cxac/liquid/websocket/private/auth/event"
	"github.com/ymiz/go-cxac/liquid/websocket/private/subscription"
	private2 "github.com/ymiz/go-cxac/liquid/websocket/private/subscription/data"
	"github.com/ymiz/go-cxac/liquid/websocket/private/subscription/data/channel/executions"
	"github.com/ymiz/go-cxac/liquid/websocket/private/subscription/data/event"
	event3 "github.com/ymiz/go-cxac/liquid/websocket/private/subscription/event"
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
	token            *common.Token
}

func NewClient(
	ch chan *Execution,
	currencyPairCode pair.Code,
	errorHandler *ErrorHandler,
	token *common.Token,
) *Client {
	return &Client{ch: ch, currencyPairCode: currencyPairCode, errorHandler: errorHandler, token: token}
}

func (c *Client) Connect() (r *result.Result) {
	connector := connection.NewConnector(private.EndPoint, nil, func(
		message []byte,
		conn *websocket.Conn,
	) *result.Result {
		ev, err := jsonparser.GetString(message, "event")
		if err != nil {
			c.errorHandler.onUnMarshalError(err)
			return nil
		}

		switch ev {
		case event2.Failure.ToString():
			return result.NewResult(code.AuthError, errors.New("receive auth_failure message"))
		case event2.Success.ToString():
			err = conn.WriteJSON(subscription.Parameter{
				Event: event3.Subscribe,
				Data: private2.Data{
					Channel: executions.Generate(c.currencyPairCode),
					Event:   event.Updated,
				},
			})
			if err != nil {
				return result.NewResult(code.Disconnected, err)
			}
			return nil
		case event3.ConnectionEstablished.ToString():
			path := "/realtime"
			headers, err := headers3.GenerateHeaders(c.token, path)
			if err != nil {
				return result.NewResult(code.AuthError, err)
			}
			err = conn.WriteJSON(auth.Parameter{
				Event: event2.Request,
				Data: data.Data{
					Headers: headers,
					Path:    path,
				},
			})
			if err != nil {
				return result.NewResult(code.AuthError, err)
			}
			return nil
		case event3.SubscriptionSucceeded.ToString():
			return nil
		case event3.Created.ToString():
			var raw *RawExecution
			err := json.Unmarshal(message, &raw)
			if err != nil {
				if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
					c.errorHandler.onUnMarshalError(err)
				}
				return nil
			}
			var d *RawData
			err = json.Unmarshal([]byte(raw.Data), &d)
			if err != nil {
				if c.errorHandler != nil && c.errorHandler.onUnMarshalError != nil {
					c.errorHandler.onUnMarshalError(err)
				}
				return nil
			}
			c.ch <- NewExecution(raw.Channel, convertData(d), raw.Event)
			return nil
		default:
			return nil
		}
	})
	return connector.Connect()
}

func convertData(r *RawData) Data {
	return Data{
		Id:            r.Id,
		Quantity:      r.Quantity,
		Price:         r.Price,
		TakerSide:     r.TakerSide,
		MySide:        r.MySide,
		CreatedAt:     time.Unix(int64(r.CreatedAt), 0),
		OrderId:       r.OrderId,
		ClientOrderId: r.ClientOrderId,
		Target:        r.Target,
		Timestamp:     r.Timestamp,
	}

}
