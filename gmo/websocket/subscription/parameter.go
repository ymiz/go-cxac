package subscription

import (
	. "github.com/ymiz/go-cxac/gmo/websocket/subscription/channel"
	. "github.com/ymiz/go-cxac/gmo/websocket/subscription/command"
	. "github.com/ymiz/go-cxac/gmo/websocket/subscription/option"
	. "github.com/ymiz/go-cxac/gmo/websocket/subscription/symbol"
)

// Parameter is Websocketでリアルタイム情報を購読するためのパラメータ
type Parameter struct {
	Command Command `json:"command"`
	Channel Channel `json:"channel"`
	Symbol  Symbol  `json:"symbol,omitempty"`
	Option  Option  `json:"option,omitempty"`
}
