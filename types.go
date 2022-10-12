
package eos_contract_api_client

import (
    "time"
    null "gopkg.in/guregu/null.v4"
)

// HTTP

type HTTPResponse struct {
    HTTPStatusCode int
}

func (resp *HTTPResponse) IsError() bool {
    return resp.HTTPStatusCode > 399
}

type APIError struct {
    Success null.Bool    `json:"success"`
    Message null.String  `json:"message"`
}

// Health

type ChainHealth struct {
    Status      string
    HeadBlock   int64
    HeadTime    time.Time
}

type RedisHealth struct {
    Status string `json:"status"`
}

type PostgresHealth struct {
    Status string                       `json:"status"`
    Readers []map[string]interface{}    `json:"readers"`
}

type HealthData struct {
    Version string          `json:"version"`
    Postgres PostgresHealth `json:"postgres"`
    Redis RedisHealth       `json:"redis"`
    Chain ChainHealth       `json:"chain"`
}

type Health struct {
    HTTPResponse
    Success bool
    Data HealthData
    QueryTime time.Time
}
