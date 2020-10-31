package bridgeapi

import (
	"fmt"
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
}

// Log is a group of logging items
type Log []LogEntry

type logOptions struct {
	Offset int `url:"offset"`
	Count  int `url:"count"`
}

// Log fetches the given number of logs from the bridge starting with the given offset.
func (c *Connection) Log(offset, count int) (Log, error) {
	options := logOptions{offset, count}
	var log Log
	if err := c.get(c.hashedURL("log", options), &log); err != nil {
		return nil, fmt.Errorf("could not fetch logs: %w", err)
	}
	return log, nil
}
