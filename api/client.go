/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "encoding/json"

    "github.com/valyala/fasthttp"

    "github.com/katena-chain/sdk-go-client/utils"
)

// Client is a fasthttp.Client wrapper to dialog with a JSON API.
type Client struct {
    fastHttpClient *fasthttp.Client
    apiUrl         string
}

// Client constructor.
func NewClient(apiUrl string) *Client {
    return &Client{
        fastHttpClient: &fasthttp.Client{},
        apiUrl:         apiUrl,
    }
}

// Get wraps the doRequest method to do a GET HTTP request.
func (c *Client) Get(route string, queryValues map[string]string) (*RawResponse, error) {
    return c.doRequest("GET", route, queryValues, nil)
}

// Post wraps the doRequest method to do a POST HTTP request.
func (c *Client) Post(route string, queryValues map[string]string, body interface{}) (*RawResponse, error) {
    return c.doRequest("POST", route, queryValues, body)
}

// doRequest uses the fasthttp.Client to call a distant api and returns a response.
// The body will be marshaled to JSON.
func (c *Client) doRequest(
    method string,
    route string,
    queryValues map[string]string,
    body interface{},
) (*RawResponse, error) {
    req := fasthttp.AcquireRequest()
    resp := fasthttp.AcquireResponse()
    defer func() {
        if req != nil {
            req.SetConnectionClose()
            fasthttp.ReleaseRequest(req)
        }
        if resp != nil {
            resp.SetConnectionClose()
            fasthttp.ReleaseResponse(resp)
        }
    }()

    uri, err := utils.GetUri(c.apiUrl, []string{route}, queryValues)
    if err != nil {
        return nil, err
    }
    req.SetRequestURI(uri.String())

    if body != nil {
        req.Header.SetContentType("application/json")
        marshaledBody, err := json.Marshal(body)
        if err != nil {
            return nil, err
        }
        req.SetBody(marshaledBody)
    }

    req.Header.SetMethod(method)

    err = c.fastHttpClient.Do(req, resp)
    if err != nil {
        return nil, err
    }

    return &RawResponse{
        StatusCode: resp.StatusCode(),
        Body:       resp.Body(),
    }, nil
}

// Response is a fasthttp.Response wrapper.
type RawResponse struct {
    StatusCode int
    Body       []byte
}
