
package eos_contract_api_client

import (
    "fmt"
    "strings"
    "github.com/imroc/req/v3"
    "github.com/sonh/qs"
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

func (c *Client) send(method string, path string, params interface{}) (*req.Response, error) {
    r := req.C().R()

    if params != nil {
        query, err := qs.NewEncoder().Values(params)
        if err != nil {
            return nil, err
        }
        r.SetQueryString(query.Encode())
    }

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
    }

    return resp, err
}

//  GetHealth - Fetches "/health" from API
// ---------------------------------------------------------
func (c *Client) GetHealth() (Health, error) {

    var health Health

    r, err := c.send("GET", "/health", nil)
    if err == nil {

        // Set HTTPStatusCode
        health.HTTPStatusCode = r.StatusCode

        // Parse json
        err = r.Unmarshal(&health)
    }
    return health, err
}

//  GetAsset - Fetches "/attomicaassets/v1/assets/{asset_id}" from API
// ---------------------------------------------------------
func (c *Client) GetAsset(asset_id string) (AssetResponse, error) {

    var asset AssetResponse

    r, err := c.send("GET", "/atomicassets/v1/assets/" + asset_id, nil)
    if err == nil {

        // Set HTTPStatusCode
        asset.HTTPStatusCode = r.StatusCode

        // Parse json
        err = r.Unmarshal(&asset)
    }
    return asset, err
}

//  GetAssets - Fetches "/attomicaassets/v1/assets" from API
// ---------------------------------------------------------
func (c *Client) GetAssets(params AssetsRequestParams) (AssetsResponse, error) {

    var assets AssetsResponse

    r, err := c.send("GET", "/atomicassets/v1/assets", params)
    if err == nil {

        // Set HTTPStatusCode
        assets.HTTPStatusCode = r.StatusCode

        // Parse json
        err = r.Unmarshal(&assets)
    }
    return assets, err
}
