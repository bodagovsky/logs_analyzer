package tests

import (
	"os"
	"testing"

	"github.com/bodagovsky/logs_out/src/filemanager"
	"github.com/stretchr/testify/assert"
)

func TestEmptyStream(t *testing.T) {
	file, err := os.CreateTemp("data/stream", "*.log")

	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())
	logs := []string{
		"[2025-06-24 12:34:56] [INFO] [client-42] - User 1931 successfully logged in",
		"[2025-06-24 12:35:10] [ERROR] [client-42] - Failed to connect to database: timeout after 5s",
		"[2025-06-24 12:35:40] [DEBUG] [client-42] - Received payload: {\"id\":1931,\"status\":\"ok\"}",
		"[2025-06-24 12:35:44] [INFO] [client-42] - User 1931 successfully logged out",
		"[2025-06-24 12:36:56] [INFO] [client-42] - User 1932 successfully logged in",
		"[2025-06-24 12:37:10] [ERROR] [client-42] - Failed to connect to database: timeout after 5s",
		"[2025-06-24 12:37:42] [DEBUG] [client-42] - Received payload: {\"id\":1932,\"status\":\"ok\"}",
		"[2025-06-24 12:36:56] [INFO] [client-42] - User 1932 successfully logged out",
	}
	sm := filemanager.New(file)

	var offsets []int64

	for _, log := range logs {
		offsets = append(offsets, sm.AppendLog(log))
	}

	lines := sm.GetLinesByOffsets(offsets[4:])

	assert.ElementsMatch(t, logs[4:], lines)

}

func TestNonEmptyStream(t *testing.T) {
	file, err := os.CreateTemp("data/stream", "*.log")

	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	defer file.Close()

	file.Write([]byte("[2025-06-24 12:34:56] [INFO] [client-42] - User successfully logged in"))
	file.Write([]byte("[2025-06-24 12:35:10] [ERROR] [client-42] - Failed to connect to database: timeout after 5s"))
	file.Write([]byte("[2025-06-24 12:35:40] [DEBUG] [client-42] - Received payload: {\"id\":123,\"status\":\"ok\"}))"))

	logs := []string{
		"[2025-06-24 12:35:44] [INFO] [client-42] - User 1931 successfully logged out",
		"[2025-06-24 12:36:56] [INFO] [client-42] - User 1932 successfully logged in",
		"[2025-06-24 12:37:10] [ERROR] [client-42] - Failed to connect to database: timeout after 5s",
		"[2025-06-24 12:37:42] [DEBUG] [client-42] - Received payload: {\"id\":1932,\"status\":\"ok\"}",
		"[2025-06-24 12:36:56] [INFO] [client-42] - User 1932 successfully logged out",
	}
	sm := filemanager.New(file)

	var offsets []int64

	for _, log := range logs {
		offsets = append(offsets, sm.AppendLog(log))
	}

	lines := sm.GetLinesByOffsets(offsets)
	assert.ElementsMatch(t, logs, lines)
}
