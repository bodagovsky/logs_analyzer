package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenIndex(t *testing.T) {
	ti := NewTokenIndex()
	cases := []struct {
		title    string
		request  string
		messages []string
		offsets  []int64
		expected []int64
	}{
		{
			title:    "multiple_words_second_word",
			messages: []string{"User", "logged", "in"},
			offsets:  []int64{0, 1, 2},
			request:  "logged",
			expected: []int64{1},
		},
		{
			title:    "multiple_words_first_word",
			messages: []string{"User", "logged", "in"},
			offsets:  []int64{0, 5, 25},
			request:  "User",
			expected: []int64{0},
		},
		{
			title:    "space_separated_words",
			messages: []string{"User successfully logged in"},
			offsets:  []int64{42},
			request:  "successfully",
			expected: []int64{42},
		},
		{
			title:    "space_separated_query",
			messages: []string{"User successfully logged in", "User successfully logged out"},
			offsets:  []int64{42, 63},
			request:  "in out",
			expected: []int64{42, 63},
		},
		{
			title:    "space_separated_words_multiple_messages",
			messages: []string{"User successfully logged in", "User successfully logged out"},
			offsets:  []int64{42, 63},
			request:  "successfully",
			expected: []int64{42, 63},
		},
		{
			title:    "space_separated_query_single_hit",
			messages: []string{"User successfully logged in", "User successfully logged out", "database connection error"},
			offsets:  []int64{42, 63, 128},
			request:  "database error",
			expected: []int64{128},
		},
		{
			title:    "no_match",
			messages: []string{"User successfully logged in", "User successfully logged out", "database connection error"},
			offsets:  []int64{42, 63, 128},
			request:  "fantastic sheep",
			expected: []int64{},
		},
		{
			title:    "empty_query",
			messages: []string{"User successfully logged in", "User successfully logged out", "database connection error"},
			offsets:  []int64{42, 63, 128},
			request:  "",
			expected: []int64{},
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			for i, message := range test.messages {
				ti.Index(message, test.offsets[i])
			}
			assert.ElementsMatch(t, test.expected, ti.GetOffsets(test.request))
			ti.Reset()
		})
	}
}
