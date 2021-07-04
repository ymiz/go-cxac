package sendchildorder

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/common"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/sendchildorder/request"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/sendchildorder/response"
)

type Client struct {
	commonClient *common.Client
}

func NewClient(commonClient *common.Client) *Client {
	return &Client{commonClient: commonClient}
}

func (c *Client) Do(body request.Body) (*resty.Response, *response.Body, error) {
	path := "/v1/me/sendchildorder"
	var responseBody response.Body

	resp, err := c.commonClient.Post(common.NewPostParameter(path, body, &responseBody))
	if err != nil {
		return nil, nil, err
	}

	return resp, &responseBody, nil
}
