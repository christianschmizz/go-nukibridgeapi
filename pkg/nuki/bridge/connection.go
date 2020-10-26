package bridge

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/rs/zerolog/log"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

type connection struct {
	bridgeHost string
	token      string
	scan       map[nuki.NukiID]*ScanResult
}

func ScanOnConnect() func(*connection) {
	return func(c *connection) {
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

func ConnectWithToken(bridgeHost, token string, options ...func(*connection)) (*connection, error) {
	conn := &connection{bridgeHost: bridgeHost, token: token, scan: map[nuki.NukiID]*ScanResult{}}
	for _, opt := range options {
		opt(conn)
	}
	return conn, nil
}

func (c *connection) hashedURL(p string, queryParams interface{}) string {
	ts := time.Now().UTC().Format(time.RFC3339)
	rnr := rand.Intn(1000)

	// Generate the hash
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s,%d,%s", ts, rnr, c.token)))
	hash := hex.EncodeToString(h.Sum(nil))

	u := url.URL{Scheme: "http", Host: c.bridgeHost, Path: p}

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

func (c *connection) get(url string, o interface{}) error {
	log.Debug().Str("url", url).Msg("")
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

func (c *connection) isKnown(nukiID nuki.NukiID) (*ScanResult, bool) {
	if len(c.scan) == 0 {
		return nil, false
	}

	r, exists := c.scan[nukiID]
	if !exists {
		return nil, false
	}
	return r, true
}
