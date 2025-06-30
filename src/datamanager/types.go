package datamanager

type Severity int

const (
	DEBUG Severity = iota
	WARNING
	ERROR
	CRITICAL
)

type Query struct {
	Text string
	From, To int64
}

type LogEntry struct {
	ClientID string
	Severity Severity
	Text string
	Timestamp int64
}

func (l LogEntry) Format () string {
	return "[2025-06-24 12:34:56] [INFO]  [client-42] - User successfully logged in\n"
}