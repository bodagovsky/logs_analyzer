package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {

	cases := []struct {
		title    string
		keys     []int64
		target   int64
		upto     bool
		expected int
	}{
		{
			title:    "simple_search",
			keys:     []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   2,
			expected: 1,
		},
		{
			title:    "target_exists_beginning",
			keys:     []int64{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			title:    "target_exists_middle",
			keys:     []int64{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			title:    "target_exists_end",
			keys:     []int64{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},

		// Target not found
		{
			title:    "target_absent_less_than_all",
			keys:     []int64{10, 20, 30},
			target:   5,
			expected: 0,
		},
		{
			title:    "target_absent_between_elements",
			keys:     []int64{10, 20, 30},
			target:   25,
			expected: 1,
		},
		{
			title:    "target_absent_greater_than_all",
			keys:     []int64{10, 20, 30},
			target:   35,
			expected: 2,
		},

		{
			title:    "between_first_and_second",
			keys:     []int64{10, 20, 30},
			target:   15,
			expected: 0,
		},

		// Edge cases
		{
			title:    "empty_array",
			keys:     []int64{},
			target:   5,
			expected: -1,
		},
		{
			title:    "single_element_match",
			keys:     []int64{7},
			target:   7,
			expected: 0,
		},
		{
			title:    "single_element_no_match_less",
			keys:     []int64{7},
			target:   6,
			expected: 0,
		},
		{
			title:    "single_element_no_match_greater",
			keys:     []int64{7},
			target:   8,
			expected: 0,
		},
		{
			title:    "duplicates_target_exists",
			keys:     []int64{5, 5, 5, 5, 5},
			target:   5,
			expected: 0, // First occurrence
		},
		{
			title:    "multiple_duplicates",
			keys:     []int64{5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8},
			target:   5,
			expected: 0,
		},
		{
			title:    "multiple_duplicates_middle",
			keys:     []int64{5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8},
			target:   6,
			expected: 5,
		},
		{
			title:    "between_duplicates",
			keys:     []int64{5, 5, 10, 10, 15},
			target:   12,
			expected: 3, // index of last 10
		},

		//more duplicates cases
		{
			title:    "duplicates_exact_match_first_occurrence",
			keys:     []int64{1, 2, 5, 5, 5, 6, 7},
			target:   5,
			expected: 2, // first 5
		},
		{
			title:    "duplicates_match_at_start",
			keys:     []int64{5, 5, 5, 6, 7},
			target:   5,
			expected: 0,
		},
		{
			title:    "duplicates_match_at_end",
			keys:     []int64{1, 2, 3, 5, 5, 5},
			target:   5,
			expected: 3,
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.expected, Binary(test.keys, test.target, 0, len(test.keys)))
		})
	}
}
