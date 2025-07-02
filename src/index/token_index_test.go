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
		expected int64
	}{
		{
			title:    "multiple_words_second_word",
			messages: []string{"User", "logged", "in"},
			offsets:  []int64{0, 1, 2},
			request:  "logged",
			expected: 1,
		},
		{
			title:    "multiple_words_first_word",
			messages: []string{"User", "logged", "in"},
			offsets:  []int64{0, 5, 25},
			request:  "User",
			expected: 0,
		},
		{
			title:    "space_separated_words",
			messages: []string{"User successfully logged in"},
			offsets:  []int64{42},
			request:  "successfully",
			expected: 42,
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			for i, message := range test.messages {
				ti.Index(message, test.offsets[i])
			}
			assert.Equal(t, test.expected, ti.GetOffsets(test.request)[0])
			ti.Reset()
		})
	}
}
