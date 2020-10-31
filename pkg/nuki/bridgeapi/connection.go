package bridgeapi

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/rs/zerolog/log"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// Connection holds all information required for communication with a bridge
type Connection struct {
	bridgeHost string
	token      string
	scan       map[nuki.ID]*ScanResult
}

// ScanOnConnect may be used as an options when connection and requests
// scanning info from the bridge on creation of the connection.
func ScanOnConnect() func(*Connection) {
	return func(c *Connection) {
		info, err := c.Info()
		if err != nil {
			panic(err)
		}
		// Cache scan results
		for _, r := range info.ScanResults {
			c.scan[*r.NukiID()] = &r
		}
	}
}

// IsValidBridgeHost checks for validity of foven address
func IsValidBridgeHost(bridgeHost string) (bool, error) {
	ip, _, err := net.SplitHostPort(bridgeHost)
	if err != nil {
		return false, fmt.Errorf("invalid host: %w", err)
	}

	if net.ParseIP(ip) == nil {
		return false, fmt.Errorf("invalid ip address: %s", ip)
	}

	return true, nil
}

// ConnectWithToken sets up a connection to the bridge using the given token for authentication
func ConnectWithToken(bridgeHost, token string, options ...func(*Connection)) (*Connection, error) {
	if ok, err := IsValidBridgeHost(bridgeHost); !ok {
		return nil, fmt.Errorf("invalid bridge host %s: %w", bridgeHost, err)
	}

	conn := &Connection{bridgeHost: bridgeHost, token: token, scan: map[nuki.ID]*ScanResult{}}
	for _, opt := range options {
		opt(conn)
	}
	return conn, nil
}

// hashedURL generates a hashed URL
func (c *Connection) hashedURL(path string, queryParams interface{}) string {
	ts := time.Now().UTC().Format(time.RFC3339)
	rnr := rand.Intn(1000)

	// Generate the hash
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s,%d,%s", ts, rnr, c.token)))
	hash := hex.EncodeToString(h.Sum(nil))

	u := url.URL{Scheme: "http", Host: c.bridgeHost, Path: path}

	// Put all the details at the query string and re-add it to the URL
	q := u.Query()
	q.Set("ts", ts)
	q.Set("rnr", strconv.Itoa(rnr))
	q.Set("hash", hash)
	u.RawQuery = q.Encode()

	if queryParams != nil && !reflect.DeepEqual(queryParams, reflect.Zero(reflect.TypeOf(queryParams)).Interface()) {
		queryString, err := query.Values(queryParams)
		if err == nil {
			u.RawQuery += "&" + queryString.Encode()
		}
	}

	return u.String()
}

func (c *Connection) get(url string, o interface{}) error {
	log.Debug().Str("url", url).Msg("calling bridge API")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&o); err != nil {
		return fmt.Errorf("could not decode response: %w\n\n%+v", err, resp.Body)
	}
	return nil
}

func (c *Connection) isKnown(nukiID nuki.ID) (*ScanResult, bool) {
	if len(c.scan) == 0 {
		return nil, false
	}

	r, exists := c.scan[nukiID]
	if !exists {
		return nil, false
	}
	return r, true
}
