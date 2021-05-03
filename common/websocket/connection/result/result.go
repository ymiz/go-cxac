package result

import . "github.com/ymiz/go-cxac/common/websocket/connection/result/code"

type Result struct {
	code  Code
	error error
}

func (r Result) Code() Code {
	return r.code
}

func (r Result) Error() error {
	return r.error
}

func NewResult(code Code, error error) *Result {
	return &Result{code: code, error: error}
}
