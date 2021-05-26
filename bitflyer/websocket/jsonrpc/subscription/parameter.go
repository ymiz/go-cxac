package subscription

type Parameter struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

type Params struct {
	Channel string `json:"channel"`
}
