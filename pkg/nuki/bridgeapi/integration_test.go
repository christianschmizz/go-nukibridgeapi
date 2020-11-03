// +build integration

package bridgeapi_test

import (
	"flag"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

var (
	host  string
	token string
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	flag.StringVar(&host, "host", "", "")
	flag.StringVar(&token, "token", "", "")
}

func bridgeConn(t *testing.T) *bridgeapi.Connection {
	conn, err := bridgeapi.ConnectWithToken(host, token)
	assert.NoError(t, err)
	return conn
}
