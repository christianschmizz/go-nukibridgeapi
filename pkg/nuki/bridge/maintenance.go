package bridge

import (
	"fmt"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	ID        string    `json:"nukiId,omitempty"`
	PairIndex int       `json:"pairIndex,omitempty"`
	BleHandle string    `json:"bleHandle,omitempty"`
	MacAddr   string    `json:"macAddr,omitempty"`
	Bytes     int       `json:"bytes,omitempty"`
}

type Log []LogEntry

type logOptions struct {
	Offset int `url:"offset"`
	Count  int `url:"count"`
}

func (c *Connection) Log(offset, count int) (Log, error) {
	options := logOptions{offset, count}
	var log Log
	if err := c.get(c.hashedURL("log", options), &log); err != nil {
		return nil, fmt.Errorf("could not fetch logs: %w", err)
	}
	return log, nil
}
