package tests

import (
	"testing"

	"github.com/bodagovsky/logs_out/src/filemanager"
	"github.com/bodagovsky/logs_out/src/index"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	logStorage := filemanager.NewLogStorage(index.NewTokenIndex(), index.NewIndexer(), index.NewTimestampIndex())

	entries := []struct {
		msg        string
		ts, offset int64
	}{
		{
			msg:    "user 1234 successfully logged in",
			ts:     1750820538,
			offset: 0,
		},
		{
			msg:    "failed to connect to database: timeout after 5s",
			ts:     1750820539,
			offset: 62,
		},
		{
			msg:    "received payload: {\"id\":1234,\"status\":\"ok\"}",
			ts:     1750820540,
			offset: 145,
		},
		{
			msg:    "user 1234 successfully logged out",
			ts:     1750820541,
			offset: 223,
		},

		// past one hour

		{
			msg:    "user 12345 successfully logged in",
			ts:     1750824138,
			offset: 245,
		},
		{
			msg:    "failed to connect to database: timeout after 5s",
			ts:     1750824139,
			offset: 275,
		},
		{
			msg:    "received payload: {\"id\":12345,\"status\":\"ok\"}",
			ts:     1750824140,
			offset: 333,
		},
		{
			msg:    "user 12345 successfully logged out",
			ts:     1750824141,
			offset: 375,
		},
	}
	for _, entry := range entries {
		logStorage.Index(entry.msg, entry.ts, entry.offset)
	}

	queries := []struct {
		msg       string
		timestamp int64
		expected  []int64
	}{
		{
			msg:       "logged in",
			timestamp: 1750820500,
			expected:  []int64{0, 223, 245, 375},
		},
		{
			msg:       "payload",
			timestamp: 1750820539,
			expected:  []int64{145, 333},
		},
		{
			msg:       "database",
			timestamp: 1750820540,
			expected:  []int64{62, 275},
		},
		{
			msg:       "user",
			timestamp: 1750820538,
			expected:  []int64{0, 223, 245, 375},
		},
		{
			msg:       "user",
			timestamp: 1750824138,
			expected:  []int64{245, 375},
		},
		{
			msg:       "ailed to connect to database",
			timestamp: 0,
			expected:  []int64{62, 275},
		},
		{
			msg:       "in out payload timeout",
			timestamp: 1750824138,
			expected:  []int64{245, 275, 333, 375},
		},
		{
			msg:       "",
			timestamp: 1750820538,
			expected:  []int64{0},
		},
	}

	for _, q := range queries {
		assert.ElementsMatch(t, q.expected, logStorage.Search(q.msg, q.timestamp))
	}

}
