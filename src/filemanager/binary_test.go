package filemanager

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
			upto:     false,
			expected: 1,
		},
		{
			title:    "target_exists_beginning",
			keys:     []int64{1, 2, 3, 4, 5},
			target:   1,
			upto:     false,
			expected: 0,
		},
		{
			title:    "target_exists_middle",
			keys:     []int64{1, 2, 3, 4, 5},
			target:   3,
			upto:     false,
			expected: 2,
		},
		{
			title:    "target_exists_end",
			keys:     []int64{1, 2, 3, 4, 5},
			target:   5,
			upto:     false,
			expected: 4,
		},

		// Target not found, upto = false
		{
			title:    "target_absent_less_than_all",
			keys:     []int64{10, 20, 30},
			target:   5,
			upto:     false,
			expected: 0,
		},
		{
			title:    "target_absent_between_elements",
			keys:     []int64{10, 20, 30},
			target:   25,
			upto:     false,
			expected: 2,
		},
		{
			title:    "target_absent_greater_than_all",
			keys:     []int64{10, 20, 30},
			target:   35,
			upto:     false,
			expected: -1,
		},

		// Target not found, upto = true
		{
			title:    "upto_less_than_all",
			keys:     []int64{10, 20, 30},
			target:   5,
			upto:     true,
			expected: -1,
		},
		{
			title:    "upto_between_first_and_second",
			keys:     []int64{10, 20, 30},
			target:   15,
			upto:     true,
			expected: 0,
		},
		{
			title:    "upto_between_second_and_third",
			keys:     []int64{10, 20, 30},
			target:   25,
			upto:     true,
			expected: 1,
		},
		{
			title:    "upto_greater_than_all",
			keys:     []int64{10, 20, 30},
			target:   35,
			upto:     true,
			expected: 2,
		},

		// Edge cases
		{
			title:    "empty_array",
			keys:     []int64{},
			target:   5,
			upto:     false,
			expected: -1,
		},
		{
			title:    "single_element_match",
			keys:     []int64{7},
			target:   7,
			upto:     false,
			expected: 0,
		},
		{
			title:    "single_element_no_match_upto_true",
			keys:     []int64{7},
			target:   6,
			upto:     true,
			expected: -1,
		},
		{
			title:    "single_element_no_match_upto_true2",
			keys:     []int64{7},
			target:   8,
			upto:     true,
			expected: 0,
		},
		{
			title:    "duplicates_target_exists",
			keys:     []int64{5, 5, 5, 5, 5},
			target:   5,
			upto:     false,
			expected: 0, // Accept any valid index — assuming you want first occurrence
		},
		{
			title:    "multiple_duplicates",
			keys:     []int64{5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8},
			target:   5,
			upto:     false,
			expected: 0,
		},
		{
			title:    "multiple_duplicates_with_upto",
			keys:     []int64{5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8},
			target:   6,
			upto:     true,
			expected: 5,
		},
		{
			title:    "duplicates_upto_between_duplicates",
			keys:     []int64{5, 5, 10, 10, 15},
			target:   12,
			upto:     true,
			expected: 3, // index of last 10
		},

		//more duplicates cases
		{
			title:    "duplicates_exact_match_first_occurrence",
			keys:     []int64{1, 2, 5, 5, 5, 6, 7},
			target:   5,
			upto:     false,
			expected: 2, // first 5
		},
		{
			title:    "duplicates_match_at_start",
			keys:     []int64{5, 5, 5, 6, 7},
			target:   5,
			upto:     false,
			expected: 0,
		},
		{
			title:    "duplicates_match_at_end",
			keys:     []int64{1, 2, 3, 5, 5, 5},
			target:   5,
			upto:     false,
			expected: 3,
		},
		{
			title:    "upto_before_duplicate_cluster",
			keys:     []int64{1, 2, 3, 5, 5, 5, 8, 9},
			target:   5,
			upto:     true,
			expected: 3, 
		},
		{
			title:    "upto_between_duplicate_clusters",
			keys:     []int64{1, 2, 4, 4, 4, 6, 6, 6, 9},
			target:   5,
			upto:     true,
			expected: 4, // last 4
		},
		{
			title:    "upto_after_duplicate_cluster",
			keys:     []int64{3, 3, 3, 5, 5, 5, 7},
			target:   6,
			upto:     true,
			expected: 5, // last 5
		},
		{
			title:    "target_with_gaps_between_clusters",
			keys:     []int64{1, 2, 3, 5, 5, 10, 10, 10, 15},
			target:   10,
			upto:     false,
			expected: 5, // first 10
		},
		{
			title:    "upto_across_multiple_duplicate_ranges",
			keys:     []int64{1, 1, 1, 3, 3, 3, 6, 6, 6},
			target:   5,
			upto:     true,
			expected: 5, // last 3
		},
		{
			title:    "target_after_many_duplicates",
			keys:     []int64{2, 2, 2, 2, 2, 2, 9},
			target:   9,
			upto:     false,
			expected: 6,
		},
		{
			title:    "target_within_sparse_duplicate_sections",
			keys:     []int64{1, 1, 3, 3, 3, 7, 9, 9, 9},
			target:   3,
			upto:     false,
			expected: 2,
		},
		{
			title:    "target_just_between_duplicate_sections",
			keys:     []int64{1, 1, 3, 3, 3, 7, 7, 9},
			target:   5,
			upto:     true,
			expected: 4, // last 3
		},
		{
			title:    "upto_between_identical_groups",
			keys:     []int64{1, 1, 1, 5, 5, 5, 10, 10, 10},
			target:   8,
			upto:     true,
			expected: 5, // last 5
		},
		{
			title:    "target_matches_entire_array",
			keys:     []int64{7, 7, 7, 7, 7},
			target:   7,
			upto:     false,
			expected: 0,
		},
		{
			title:    "upto_no_element_less_than_target_with_duplicates",
			keys:     []int64{5, 5, 5},
			target:   3,
			upto:     true,
			expected: -1,
		},
		{
			title:    "target_in_middle_of_duplicates",
			keys:     []int64{1, 2, 5, 5, 5, 7, 8},
			target:   5,
			upto:     false,
			expected: 2,
		},
		{
			title:    "upto_with_duplicates_at_start",
			keys:     []int64{1, 1, 2, 3, 4, 5},
			target:   2,
			upto:     true,
			expected: 2, // last 1
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.expected, Binary(test.keys, test.target, 0, len(test.keys), test.upto))
		})
	}
}
