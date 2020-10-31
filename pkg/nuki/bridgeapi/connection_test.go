package bridgeapi_test

import (
	"fmt"
	"testing"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

func ExampleConnectWithToken() {
	conn, err := bridgeapi.ConnectWithToken("192.168.1.11:8080", "abcdef")
	if err != nil {
		panic(err)
	}
	info, err := conn.Info()
	if err != nil {
		panic(err)
	}
	for _, result := range info.ScanResults {
		fmt.Println(result.Name)
	}
}

func ExampleScanOnConnect() {
	_, err := bridgeapi.ConnectWithToken("192.168.1.11:8080", "abcdef", bridgeapi.ScanOnConnect())
	if err != nil {
		panic(err)
	}
}

func TestIsValidBridgeHost(t *testing.T) {
	tests := []struct {
		host    string
		isValid bool
	}{
		{"test 2", false},
		{"192.168.1.1", false},
		{"192.168.1.1.11", false},
		{"192.168.1.1:8080", true},
		{"2001:db8::68", false},
		{"2001:0db8:0000:08d3:0000:8a2e:0070:7344", false},
		{"2001:db8:0:8d3:0:8a2e:70:7344", false},
		{"[2001:db8::68]", false},
		{"[2001:0db8:0000:08d3:0000:8a2e:0070:7344]", false},
		{"[2001:db8:0:8d3:0:8a2e:70:7344]:8080", true},
		{"https://192.168.1.1", false},
		{"https://1", false},
		{"http://[::ffff:192.0.2.1]:8080", false},
		{"https://192.168.1.1/fail", false},
	}
	for _, tt := range tests {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.host, func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other
			if ok, err := bridgeapi.IsValidBridgeHost(tt.host); ok != tt.isValid {
				t.Errorf("expected \"%s\" to be %v: %s", tt.host, tt.isValid, err)
			}
		})
	}
}
