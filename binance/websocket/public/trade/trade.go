package trade

// Trade is https://binance-docs.github.io/apidocs/spot/en/#trade-streams
type Trade struct {
	E  string `json:"e"`
	E1 int    `json:"E"`
	S  string `json:"s"`
	T  int    `json:"t"`
	P  string `json:"p"`
	Q  string `json:"q"`
	B  int    `json:"b"`
	A  int    `json:"a"`
	T1 int    `json:"T"`
	M  bool   `json:"m"`
	M1 bool   `json:"M"`
}
