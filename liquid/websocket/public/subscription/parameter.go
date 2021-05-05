package subscription

import "github.com/ymiz/go-cxac/liquid/websocket/public/subscription/event"
import event2 "github.com/ymiz/go-cxac/liquid/websocket/public/subscription/data/event"

type Parameter struct {
	Event event.Event `json:"event"`
	Data  Data        `json:"data"`
}

type Data struct {
	Channel string       `json:"channel"`
	Event   event2.Event `json:"event"`
}
