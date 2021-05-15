package subscription

import (
	"github.com/ymiz/go-cxac/liquid/websocket/private/subscription/data"
	"github.com/ymiz/go-cxac/liquid/websocket/private/subscription/event"
)

type Parameter struct {
	Event event.Event  `json:"event"`
	Data  private.Data `json:"data"`
}
