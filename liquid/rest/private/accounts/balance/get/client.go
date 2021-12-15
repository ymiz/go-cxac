package get

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/liquid/rest/private/accounts/balance/get/response"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
)

type Client struct {
	token *common.Token
	rc    *resty.Client
}

func NewClient(token *common.Token, rc *resty.Client) *Client {
	return &Client{token: token, rc: rc}
}

func (c *Client) Do() (*resty.Response, *response.Body, error) {
	path := "/accounts/balance/"

	cl, err := common.GenerateRequest(c.rc, c.token, path)
	if err != nil {
		return nil, nil, err
	}

	resp, err := cl.Get(common.EndPoint + path)
	if err != nil {
		return nil, nil, err
	}

	// readObjectStart: expect { or n, but found [, error found in #1 byte of... のエラーが出るので。配列に名前を付けておかないとうまくunmarshalできない。
	b := []byte("{\"models\":" + string(resp.Body()) + "}")

	var responseBody *response.Body
	err = json.Unmarshal(b, &responseBody)
	if err != nil {
		return nil, nil, err
	}

	return resp, responseBody, nil
}
