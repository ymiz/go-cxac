package auth

import (
	"github.com/ymiz/go-cxac/liquid/websocket/private/auth/data"
	"github.com/ymiz/go-cxac/liquid/websocket/private/auth/event"
)

type Parameter struct {
	Event event.Event `json:"event"`
	Data  data.Data   `json:"data"`
}
