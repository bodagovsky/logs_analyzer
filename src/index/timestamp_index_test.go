package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimestampIndex(t *testing.T) {
	tsi := NewTimestampIndex()

	cases := []struct {
		title      string
		logEntries [][2]int64
		query      int64
		expected   int64
	}{
		{
			title: "exact_match",
			logEntries: [][2]int64{
				{1_750_468_320, 0},
				{1_750_468_321, 1},
				{1_750_468_322, 2},
				{1_750_468_323, 3},
				{1_750_468_324, 4},
				// hour
				{1_750_475_320, 5},
				{1_750_475_321, 6},
				{1_750_475_322, 7},
				{1_750_475_323, 8},
				{1_750_475_324, 9},
				// hour
				{1_750_480_320, 10},
				{1_750_480_321, 11},
				{1_750_480_322, 12},
				{1_750_480_323, 13},
				{1_750_480_324, 14},
			},
			query:    1_750_468_320,
			expected: 0,
		},
		{
			title: "exact_match_next_hour",
			logEntries: [][2]int64{
				{1_750_468_320, 0},
				{1_750_468_321, 1},
				{1_750_468_322, 2},
				{1_750_468_323, 3},
				{1_750_468_324, 4},
				// hour
				{1_750_475_320, 5},
				{1_750_475_321, 6},
				{1_750_475_322, 7},
				{1_750_475_323, 8},
				{1_750_475_324, 9},
				// hour
				{1_750_480_320, 10},
				{1_750_480_321, 11},
				{1_750_480_322, 12},
				{1_750_480_323, 13},
				{1_750_480_324, 14},
			},
			query:    1_750_475_320,
			expected: 5,
		},
		{
			title: "exact_match_last_hour",
			logEntries: [][2]int64{
				{1_750_468_320, 0},
				{1_750_468_321, 1},
				{1_750_468_322, 2},
				{1_750_468_323, 3},
				{1_750_468_324, 4},
				// hour
				{1_750_475_320, 5},
				{1_750_475_321, 6},
				{1_750_475_322, 7},
				{1_750_475_323, 8},
				{1_750_475_324, 9},
				// hour
				{1_750_480_320, 10},
				{1_750_480_321, 11},
				{1_750_480_322, 12},
				{1_750_480_323, 13},
				{1_750_480_324, 14},
			},
			query:    1_750_480_320,
			expected: 10,
		},
		{
			title: "earliest_entry",
			logEntries: [][2]int64{
				{1_750_468_320, 0},
				{1_750_468_321, 1},
				{1_750_468_322, 2},
				{1_750_468_323, 3},
				{1_750_468_324, 4},
				// hour
				{1_750_475_320, 5},
				{1_750_475_321, 6},
				{1_750_475_322, 7},
				{1_750_475_323, 8},
				{1_750_475_324, 9},
				// hour
				{1_750_480_320, 10},
				{1_750_480_321, 11},
				{1_750_480_322, 12},
				{1_750_480_323, 13},
				{1_750_480_324, 14},
			},
			query:    1_750_468_319,
			expected: 0,
		},
		{
			title: "latest_entry",
			logEntries: [][2]int64{
				{1_750_468_320, 0},
				{1_750_468_321, 1},
				{1_750_468_322, 2},
				{1_750_468_323, 3},
				{1_750_468_324, 4},
				// hour
				{1_750_475_320, 5},
				{1_750_475_321, 6},
				{1_750_475_322, 7},
				{1_750_475_323, 8},
				{1_750_475_324, 9},
				// hour
				{1_750_480_320, 10},
				{1_750_480_321, 11},
				{1_750_480_322, 12},
				{1_750_480_323, 13},
				{1_750_480_324, 14},
			},
			query:    1_750_480_325,
			expected: 10,
		},
		{
			title: "between_hours",
			logEntries: [][2]int64{
				{1_750_468_320, 0},
				{1_750_468_321, 1},
				{1_750_468_322, 2},
				{1_750_468_323, 3},
				{1_750_468_324, 4},
				// hour
				{1_750_475_320, 5},
				{1_750_475_321, 6},
				{1_750_475_322, 7},
				{1_750_475_323, 8},
				{1_750_475_324, 9},
				// hour
				{1_750_480_320, 10},
				{1_750_480_321, 11},
				{1_750_480_322, 12},
				{1_750_480_323, 13},
				{1_750_480_324, 14},
			},
			query:    1_750_476_324,
			expected: 5,
		},
		{
			title: "between_hours_2",
			logEntries: [][2]int64{
				{1_750_468_320, 0},
				{1_750_468_321, 1},
				{1_750_468_322, 2},
				{1_750_468_323, 3},
				{1_750_468_324, 4},
				// hour
				{1_750_475_320, 5},
				{1_750_475_321, 6},
				{1_750_475_322, 7},
				{1_750_475_323, 8},
				{1_750_475_324, 9},
				// hour
				{1_750_480_320, 10},
				{1_750_480_321, 11},
				{1_750_480_322, 12},
				{1_750_480_323, 13},
				{1_750_480_324, 14},
			},
			query:    1_750_469_324,
			expected: 0,
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			for _, entry := range test.logEntries {
				tsi.InsertLogEntry(entry[0], entry[1])
			}
			assert.Equal(t, test.expected, tsi.LocateLogEntry(test.query))
			tsi.Reset()
		})
	}

}
