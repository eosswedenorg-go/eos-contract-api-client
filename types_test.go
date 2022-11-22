package eos_contract_api_client

import (
    "time"

    "reflect"
    "testing"
)

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

func TestUnixTime_Time(t *testing.T) {
    tests := []struct {
        name string
        ts   UnixTime
        want time.Time
    }{
        { "Epoc", UnixTime(0), time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC) },
        { "Two Seconds after epoch", UnixTime(2000), time.Date(1970, time.January, 1, 0, 0, 2, 0, time.UTC) },
        { "Date1", UnixTime(1644612684432), time.Date(2022, time.February, 11, 20, 51, 24, 432, time.UTC) },
        { "Date2", UnixTime(1831324037241), time.Date(2028, time.January, 12, 21, 07, 17, 241, time.UTC) },
        { "Date3", UnixTime(1272908563433), time.Date(2010, time.May, 3, 17, 42, 43, 433, time.UTC) },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.ts.Time(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("UnixTime.Time() = %v, want %v", got, tt.want)
            }
        })
    }
}
