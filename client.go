
package eos_contract_api_client

import (
    "fmt"
    "strings"
    "errors"
    "github.com/imroc/req/v3"
)

type Client struct {
    Url string
    Host string
}

func New(url string) (*Client) {
    return &Client{
        Url: url,
    }
}

func isContentType(t string, expected string) bool {

    p := strings.IndexByte(t, ';')
    if p >= 0 {
        t = t[:p]
    }
    return t == expected
}

func (c *Client) send(method string, path string) (*req.Response, error) {
    r := req.C().R()

    if len(c.Host) > 0 {
        r.SetHeader("Host", c.Host)
    }

    resp, err := r.Send(method, c.Url + path)
    if err != nil {
        return nil, err
    }

    t := resp.GetContentType()
    if ! isContentType(t, "application/json") {
        return nil, fmt.Errorf("invalid content-type '%s', expected 'application/json'", t)
    }


    if resp.IsError() {
        r_err := APIError{}
        if resp.Unmarshal(&r_err) == nil && r_err.Success.Valid && !r_err.Success.Bool {
            return nil, fmt.Errorf("API Error: %s", r_err.Message.String)
        }
        return nil, errors.New(resp.Status)
    }

    return resp, err
}

//  GetHealth - Fetches "/health" from API
// ---------------------------------------------------------
func (c *Client) GetHealth() (Health, error) {

    var health Health

    r, err := c.send("GET", "/health")
    if err == nil {

        // Set HTTPStatusCode
        health.HTTPStatusCode = r.StatusCode

        // Parse json
        err = r.Unmarshal(&health)
    }
    return health, err
}
