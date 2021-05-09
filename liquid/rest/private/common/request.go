package common

import "github.com/go-resty/resty/v2"

func GenerateRequest(client *resty.Client, token *Token, path string) (*resty.Request, error) {
	sign, err := GenerateSignature(token, path)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"content-type":         "application/json; charset=UTF-8",
		"Accept-Encoding":      "gzip",
		"X-Quoine-API-Version": "2",
		"X-Quoine-Auth":        sign,
	}
	cl := client.R().SetHeaders(header)
	return cl, nil
}
