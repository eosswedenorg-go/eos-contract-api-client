
package eos_contract_api_client

import (
    "time"
    "encoding/json"
)

func (c *ChainHealth) UnmarshalJSON(s []byte) error {

    var r struct {
        S string  `json:"status"`
        B int64   `json:"head_block"`
        T int64   `json:"head_time"`
    }

	err := json.Unmarshal(s, &r)
	if err != nil {
		return err
	}

    c.Status = r.S
    c.HeadBlock = r.B
    c.HeadTime = fromTS(r.T)

	return nil
}

func (h *Health) UnmarshalJSON(s []byte) error {

    var r struct {
        S bool       `json:"success"`
        D HealthData `json:"data"`
        T int64      `json:"query_time"`
    }

	err := json.Unmarshal(s, &r)
	if err != nil {
		return err
	}

    h.Success = r.S
    h.Data = r.D
    h.QueryTime = fromTS(r.T)

	return nil
}

func (a *AssetsResponse) UnmarshalJSON(s []byte) error {

    var r struct {
        S bool       `json:"success"`
        D []Asset    `json:"data"`
        T int64      `json:"query_time"`
    }

	err := json.Unmarshal(s, &r)
	if err != nil {
		return err
	}

    a.Success = r.S
    a.Data = r.D
    a.QueryTime = fromTS(r.T)

	return nil
}

func fromTS(ts int64) time.Time {
    return time.Unix(ts / 1000, ts % 1000).UTC()
}
