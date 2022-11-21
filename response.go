
package eos_contract_api_client

import (
    null "gopkg.in/guregu/null.v4"
)

// HTTP

type HTTPResponse struct {
    HTTPStatusCode int
}

func (resp *HTTPResponse) IsError() bool {
    return resp.HTTPStatusCode == 0 || resp.HTTPStatusCode > 399
}

// API

type APIError struct {
    Success null.Bool    `json:"success"`
    Message null.String  `json:"message"`
}

type APIResponse struct {
    HTTPResponse
    Success bool        `json:"success"`
    QueryTime UnixTime  `json:"query_time"`
}

// Health
type Health struct {
    APIResponse
    Data HealthData
}

// Assets

type AssetResponse struct {
    APIResponse
    Data Asset
}

type AssetsResponse struct {
    APIResponse
    Data []Asset
}