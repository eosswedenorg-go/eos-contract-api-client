package eos_contract_api_client

import "testing"

func TestHTTPResponse_IsError(t *testing.T) {
    tests := []struct {
        name   string
        code   int
        want   bool
    }{
        {"0 code is error", 0, true},
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

func TestUnixTime_UnmarshalJson(t *testing.T) {

    tests := []struct {
        name   string
        input []byte
        expectErr bool
        expected UnixTime
    }{
        { "number", []byte("1074932802"), false, UnixTime(1074932802) },
        { "number nanoseconds", []byte("1800718379432"), false, UnixTime(1800718379432) },
        { "string", []byte("\"1476870484\""), false, UnixTime(1476870484) },
        { "string nanoseconds", []byte("\"1440894197834\""), false, UnixTime(1440894197834) },
        { "random", []byte{0x1, 0xff, 0x3c}, true, UnixTime(0) },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

            var ts UnixTime

            if err := ts.UnmarshalJSON(tt.input); (err != nil) != tt.expectErr {
                t.Errorf("UnixTime.UnmarshalJSON(%s) expected error but got <nil>", string(tt.input))
            }

            if ts != tt.expected {
                t.Errorf("UnixTime.UnmarshalJSON(%s) parsed value = %v, expected: %v", string(tt.input), ts, tt.expected)
            }
        })
    }
}
