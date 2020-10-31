package bridgeapi_test

import (
	"fmt"

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
