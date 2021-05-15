package data

import (
	headers2 "github.com/ymiz/go-cxac/liquid/websocket/private/auth/data/headers"
)

type Data struct {
	Headers *headers2.Headers `json:"headers"`
	Path    string            `json:"path"`
}
