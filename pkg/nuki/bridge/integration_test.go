// +build integration

package bridge_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

var (
	host  = flag.String("host", "", "")
	token = flag.String("token", "", "")
)

func bridgeConn(t *testing.T, host, token string) *nukibridge.Connection {
	t.Logf("args: %s", strings.Join(flag.Args(), " "))
	t.Logf("test.timeout: %v", flag.Lookup("test.timeout").Value)

	conn, err := nukibridge.ConnectWithToken(host, token)
	assert.NoError(t, err)
	return conn
}

func init() {
	//flag.Parse()
}
