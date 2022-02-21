
package eos_contract_api_client

import (
    "time"
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestGetHealth(t *testing.T) {

    var srv = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
        if req.URL.String() == "/health" {
            payload := `{
                "success":true,
                "data":{
                    "version":"1.0.0",
                    "postgres":{
                        "status":"OK",
                        "readers":[
                            {
                                "block_num":"167836036"
                            },
                            {
                                "block_num":"167836034"
                            }
                        ]
                    },
                    "redis":{
                        "status":"OK"
                    },
                    "chain":{
                        "status":"OK",
                        "head_block":167836035,
                        "head_time":1645374771500
                    }
                },
                "query_time":1645374772067
            }`

            res.Write([]byte(payload))
        }
    }))

    client := New(srv.URL)

    h, err := client.GetHealth()

    require.NoError(t, err)
    assert.Equal(t, 200, h.HTTPStatusCode)

    assert.True(t, h.Success)
    assert.Equal(t, time.Time(time.Date(2022, time.February, 20, 16, 32, 52, 67, time.UTC)), h.QueryTime)

    // Data
    assert.Equal(t, "1.0.0", h.Data.Version)

    // Postgres
    assert.Equal(t, "OK", h.Data.Postgres.Status)

    // Redis
    assert.Equal(t, "OK", h.Data.Redis.Status)

    // Chain
    assert.Equal(t, "OK", h.Data.Chain.Status)
    assert.Equal(t, int64(167836035), h.Data.Chain.HeadBlock)
    assert.Equal(t, time.Time(time.Date(2022, time.February, 20, 16, 32, 51, 500, time.UTC)), h.Data.Chain.HeadTime)
}

func TestGetHealtPostgresFailed(t *testing.T) {

    var srv = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
        if req.URL.String() == "/health" {
            payload := `{
                "success":true,
                "data":{
                    "version":"1.0.0",
                    "postgres":{
                        "status":"ERROR",
                        "readers":[]
                    },
                    "redis":{
                        "status":"ERROR"
                    },
                    "chain":{
                        "status":"ERROR",
                        "head_block":0,
                        "head_time":0
                    }
                },
                "query_time":1645374772067
            }`

            res.Write([]byte(payload))
        }
    }))

    client := New(srv.URL)

    h, err := client.GetHealth()

    require.NoError(t, err)
    assert.Equal(t, 200, h.HTTPStatusCode)

    assert.True(t, h.Success)
    assert.Equal(t, time.Time(time.Date(2022, time.February, 20, 16, 32, 52, 67, time.UTC)), h.QueryTime)

    // Data
    assert.Equal(t, "1.0.0", h.Data.Version)

    // Postgres
    assert.Equal(t, "ERROR", h.Data.Postgres.Status)

    // Redis
    assert.Equal(t, "ERROR", h.Data.Redis.Status)

    // Chain
    assert.Equal(t, "ERROR", h.Data.Chain.Status)
    assert.Equal(t, int64(0), h.Data.Chain.HeadBlock)

    assert.Equal(t, time.Unix(0, 0).UTC(), h.Data.Chain.HeadTime)
}
