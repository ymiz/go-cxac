package cancelchildorder

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/cancelchildorder/request"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/common"
)

type Client struct {
	commonClient *common.Client
}

func NewClient(commonClient *common.Client) *Client {
	return &Client{commonClient: commonClient}
}

func (c Client) Do(body request.Body) (*resty.Response, error) {
	path := "/v1/me/cancelchildorder"

	resp, err := c.commonClient.Post(common.NewPostParameter(path, body, nil))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
