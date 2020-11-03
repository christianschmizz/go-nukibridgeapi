package bridgeapi

import (
	"net/http"
	"time"
)

// LogEntry is a single logging item from the bridge
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	ID        string    `json:"nukiId,omitempty"`
	PairIndex int       `json:"pairIndex,omitempty"`
	BleHandle string    `json:"bleHandle,omitempty"`
	MacAddr   string    `json:"macAddr,omitempty"`
	Bytes     int       `json:"bytes,omitempty"`
	Count     int       `json:"count,omitempty"`
	ServerNum int       `json:"serverNum,omitempty"`
}

// Log is a group of logging items
type Log []LogEntry

type logOptions struct {
	Offset int `url:"offset"`
	Count  int `url:"count"`
}

// Log fetches the given number of logs from the bridge starting with the given offset.
func (c *Connection) Log(offset, count int) (Log, error) {
	options := &logOptions{
		Offset: offset,
		Count:  count,
	}

	resp, err := c.request("log", options)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	}

	var data Log
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
