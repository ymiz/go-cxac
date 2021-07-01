package ticker

type Response struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"channelMessage"`
	Params  Params `json:"params"`
}

type Params struct {
	Channel string `json:"channel"`
	Message Ticker `json:"message"`
}
