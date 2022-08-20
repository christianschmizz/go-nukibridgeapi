package bridgeapi

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/crypto/nacl/secretbox"

	"github.com/google/go-querystring/query"
	"github.com/rs/zerolog/log"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

var (
	// ErrInvalidToken is issued as soon a a token is invalid or the token parameter is missing
	ErrInvalidToken = errors.New("token is invalid or a hashed token parameter is missing")

	// ErrUnknownDevice is issued when the given Nuki device is unknown
	ErrUnknownDevice = errors.New("the given Nuki device is unknown")

	// ErrDeviceIsTemporarilyOffline is issued when the Nuki device is temporarily offline
	ErrDeviceIsTemporarilyOffline = errors.New("the given Nuki device is offline")

	// ErrInvalidURL is issued when the given URL is invalid or too long
	ErrInvalidURL = errors.New("the given URL is invalid or too long")

	// ErrAnotherRequestIsAlreadyRunning is issued when the device is already busy
	ErrAnotherRequestIsAlreadyRunning = errors.New("another request already running on the device. Increase intervals between API calls sent to the Bridge as it can only handle one request at a time")

	// ErrDeviceIsOffline indicates that the device was not seen by the bridge for a longer period of time
	ErrDeviceIsOffline = errors.New("the device was noticed to be offline for a longer period of time")
)

// ErrInvalidAction is issued when the given action was invalid
type ErrInvalidAction struct {
	Action nuki.LockAction
}

func (e *ErrInvalidAction) Error() string { return fmt.Sprintf("action %d is invalid", e.Action) }

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Connection holds all information required for communication with a bridge
type Connection struct {
	client       HTTPClient
	bridgeHost   string
	token        string
	scan         map[nuki.ID]*ScanResult
	encryptToken bool
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

// UseClient uses the given client
func UseClient(client HTTPClient) func(*Connection) {
	return func(c *Connection) {
		c.client = client
	}
}

// UseEncryptedToken enables the use of an encrypted token introduces w/ 1.13.2 (17.06.2022)
func UseEncryptedToken() func(*Connection) {
	return func(c *Connection) {
		c.encryptToken = true
	}
}

// IsValidBridgeHost checks for validity of given address
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

	conn := &Connection{
		bridgeHost: bridgeHost,
		token:      token,
		scan:       map[nuki.ID]*ScanResult{},
	}
	for _, opt := range options {
		opt(conn)
	}
	if conn.client == nil {
		conn.client = &http.Client{
			Timeout: time.Duration(10) * time.Second,
		}
	}
	return conn, nil
}

// prepareURL generates a hashed URL
func (c *Connection) prepareURL(path string, queryParams interface{}) string {
	ts := time.Now().UTC().Format(time.RFC3339)
	rnr := rand.Intn(1000)

	u := url.URL{Scheme: "http", Host: c.bridgeHost, Path: path}
	q := u.Query()

	if c.encryptToken {
		nonce := [24]byte{}
		_, err := rand.Read(nonce[:])
		if err != nil {
			panic(err)
		}

		h := sha256.New()
		h.Write([]byte(c.token))
		key := [32]byte{}
		copy(key[:], h.Sum(nil))

		ctoken := secretbox.Seal(nil, []byte(ts+`,`+strconv.Itoa(rnr)), &nonce, &key)

		q.Set("ctoken", hex.EncodeToString(ctoken))
		q.Set("nonce", hex.EncodeToString(nonce[:]))
	} else {
		// Generate the hash
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("%s,%d,%s", ts, rnr, c.token)))
		hash := hex.EncodeToString(h.Sum(nil))

		q.Set("ts", ts)
		q.Set("rnr", strconv.Itoa(rnr))
		q.Set("hash", hash)
	}

	// Put all the details at the query string and re-add it to the URL
	u.RawQuery = q.Encode()

	if queryParams != nil && !reflect.DeepEqual(queryParams, reflect.Zero(reflect.TypeOf(queryParams)).Interface()) {
		queryString, err := query.Values(queryParams)
		if err == nil {
			u.RawQuery += "&" + queryString.Encode()
		}
	}

	return u.String()
}

func (c *Connection) request(path string, queryParams interface{}) (*APIResponseHandler, error) {
	requestURL := c.prepareURL(path, queryParams)
	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	log.Debug().Str("url", requestURL).Msg("requesting")
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("requesting '%s' failed: %w", path, err)
	}
	defer resp.Body.Close()

	r, err := NewAPIResponseHandler(resp)
	if err != nil {
		return nil, err
	}

	if r.Is(http.StatusGone) {
		return nil, ErrDeviceIsOffline
	}

	return r, nil
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
