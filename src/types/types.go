package requestmanager

import (
	"fmt"
	"time"
)

type Severity int

const (
	DEBUG Severity = iota
	WARNING
	ERROR
	CRITICAL
)

func (s Severity) String() string {
	switch s {
	case DEBUG:
		return "DEBUG"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	default:
		return ""
	}
}

type Query struct {
	ClientID string
	Text     string
	From, To int64
}

type LogEntry struct {
	ClientID  string
	Severity  Severity
	Text      string
	Timestamp int64
}

func (l LogEntry) Format() string {
	// TODO: change the way date is formatted
	// right now it introduces timezone skew
	t := time.Unix(l.Timestamp, 0)
	return fmt.Sprintf("[%s] [%s] [%s] - [%s]", t.Format(time.DateTime), l.Severity, l.ClientID, l.Text)
}
