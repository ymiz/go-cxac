package get

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/common/json"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"github.com/ymiz/go-cxac/liquid/rest/private/transactions/get/request/parameter"
	"github.com/ymiz/go-cxac/liquid/rest/private/transactions/get/response"
)

type Client struct {
	token *common.Token
	rc    *resty.Client
}

func NewClient(token *common.Token, rc *resty.Client) *Client {
	return &Client{token: token, rc: rc}
}

func (c *Client) Do(parameterGenerator *parameter.Generator) (*resty.Response, *response.Body, error) {
	path := "/transactions/"
	if parameterGenerator != nil {
		path = path + "?" + parameterGenerator.Generate().Encode()
	}

	cl, err := common.GenerateRequest(c.rc, c.token, path)
	if err != nil {
		return nil, nil, err
	}

	resp, err := cl.Get(common.EndPoint + path)
	if err != nil {
		return nil, nil, err
	}

	var responseBody *response.Body
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		return nil, nil, err
	}

	return resp, responseBody, nil
}
