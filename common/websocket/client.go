package websocket

import (
	. "github.com/ymiz/go-cxac/common/websocket/connection/result"
)

type Client interface {
	Connect() *Result
}
