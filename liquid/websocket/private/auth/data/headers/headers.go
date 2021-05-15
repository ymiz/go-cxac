package headers

type Headers struct {
	XQuoineAuth string `json:"X-Quoine-Auth"`
}

func NewHeaders(XQuoineAuth string) *Headers {
	return &Headers{XQuoineAuth: XQuoineAuth}
}
