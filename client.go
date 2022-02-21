
package eos_contract_api_client

import (
    "fmt"
    "strings"
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
        return nil, fmt.Errorf("Invalid content-type '%s', expected 'application/json'", t)
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
