// +build integration

package bridgeapi_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

var (
	host  = flag.String("host", "", "")
	token = flag.String("token", "", "")
)

func bridgeConn(t *testing.T, host, token string) *bridgeapi.Connection {
	t.Logf("args: %s", strings.Join(flag.Args(), " "))
	t.Logf("test.timeout: %v", flag.Lookup("test.timeout").Value)

	conn, err := bridgeapi.ConnectWithToken(host, token)
	assert.NoError(t, err)
	return conn
}
