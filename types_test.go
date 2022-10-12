package eos_contract_api_client

import "testing"

func TestHTTPResponse_IsError(t *testing.T) {
	tests := []struct {
		name   string
		code   int
		want   bool
	}{
        {"400 code is error", 400, true},
        {"400 codes is error", 404, true},
        {"500 code is error", 500, true},
        {"500 codes is error", 502, true},
        {"300 codes is not error", 312, false},
        {"200 codes is not error", 202, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &HTTPResponse{
				HTTPStatusCode: tt.code,
			}
			if got := resp.IsError(); got != tt.want {
				t.Errorf("HTTPResponse.IsError() = %v, want %v", got, tt.want)
			}
		})
	}
}
