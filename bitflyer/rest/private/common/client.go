package common

import (
	"github.com/go-resty/resty/v2"
	"github.com/ymiz/go-cxac/bitflyer/rest/private/common/authentication"
	method2 "github.com/ymiz/go-cxac/bitflyer/rest/private/common/method"
	"github.com/ymiz/go-cxac/common/json"
	"net/url"
)

type GetParameter struct {
	path             string
	requestParameter url.Values
	response         interface{}
}

func NewGetParameter(path string, requestParameter url.Values, response interface{}) *GetParameter {
	return &GetParameter{path: path, requestParameter: requestParameter, response: response}
}

type PostParameter writeParameter
type PutParameter writeParameter
type DeleteParameter writeParameter

type writeParameter struct {
	path        string
	requestBody interface{}
	response    interface{}
}

func NewPostParameter(path string, requestBody, response interface{}) *PostParameter {
	return &PostParameter{path: path, requestBody: requestBody, response: response}
}
func NewPutParameter(path string, requestBody, response interface{}) *PutParameter {
	return &PutParameter{path: path, requestBody: requestBody, response: response}
}
func NewDeleteParameter(
	path string,
	requestBody,
	response interface{},
) *DeleteParameter {
	return &DeleteParameter{path: path, requestBody: requestBody, response: response}
}

type Client struct {
	client *resty.Client
	auth   *authentication.Information
}

func NewClient(client *resty.Client, auth *authentication.Information) *Client {
	return &Client{client: client, auth: auth}
}

func (c *Client) Get(param *GetParameter) (*resty.Response, error) {
	method := method2.Get
	requestStr := ""
	if param.requestParameter != nil {
		requestStr = "?" + param.requestParameter.Encode()
	}
	header := GenerateHeader(c.auth.AccessKey(), method.ToString(), param.path, requestStr, c.auth.SecretKey())
	cl := c.client.R().SetHeaders(header)
	resp, err := cl.Get(EndPoint + param.path + requestStr)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body(), &param.response)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) Post(param *PostParameter) (*resty.Response, error) {
	method := method2.Post
	requestByte, err := json.Marshal(param.requestBody)
	if err != nil {
		return nil, err
	}
	requestStr := string(requestByte)

	header := GenerateHeader(c.auth.AccessKey(), method.ToString(), param.path, requestStr, c.auth.SecretKey())
	cl := c.client.R().SetHeaders(header)
	resp, err := cl.SetBody(requestStr).Post(EndPoint + param.path)
	if err != nil {
		return nil, err
	}

	if param.response != nil {
		err = json.Unmarshal(resp.Body(), &param.response)
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func (c *Client) Put(param *PutParameter) (*resty.Response, error) {
	method := method2.Put
	requestByte, err := json.Marshal(param.requestBody)
	if err != nil {
		return nil, err
	}
	requestStr := string(requestByte)

	header := GenerateHeader(c.auth.AccessKey(), method.ToString(), param.path, requestStr, c.auth.SecretKey())
	cl := c.client.R().SetHeaders(header)
	resp, err := cl.SetBody(requestStr).Put(EndPoint + param.path)
	if err != nil {
		return nil, err
	}

	if param.response != nil {
		err = json.Unmarshal(resp.Body(), &param.response)
		if err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (c *Client) Delete(param *PutParameter) (*resty.Response, error) {
	method := method2.Delete
	requestByte, err := json.Marshal(param.requestBody)
	if err != nil {
		return nil, err
	}
	requestStr := string(requestByte)

	header := GenerateHeader(c.auth.AccessKey(), method.ToString(), param.path, requestStr, c.auth.SecretKey())
	cl := c.client.R().SetHeaders(header)
	resp, err := cl.SetBody(requestStr).Delete(EndPoint + param.path)
	if err != nil {
		return nil, err
	}

	if param.response != nil {
		err = json.Unmarshal(resp.Body(), &param.response)
		if err != nil {
			return resp, err
		}
	}
	return resp, nil
}
