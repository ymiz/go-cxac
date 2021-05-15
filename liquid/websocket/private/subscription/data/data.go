package private

import "github.com/ymiz/go-cxac/liquid/websocket/private/subscription/data/event"

type Data struct {
	Channel string      `json:"channel"`
	Event   event.Event `json:"event"`
}
