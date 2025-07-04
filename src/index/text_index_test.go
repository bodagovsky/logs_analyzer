package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	indexer := NewIndexer()

	cases := []struct {
		title    string
		words    []string
		word     string
		expected []string
	}{
		{
			title: "simple_match",
			words: []string{"cat", "cater", "cattle", "dog", "catalog"},
			word:  "cat",
			expected: []string{
				"cat", "cater", "cattle", "catalog",
			},
		},
		{
			title: "exact_match_only",
			words: []string{"cat", "dog", "elephant"},
			word:  "dog",
			expected: []string{
				"dog",
			},
		},
		{
			title:    "no_matches",
			words:    []string{"alpha", "beta", "gamma"},
			word:     "z",
			expected: []string{},
		},
		{
			title:    "empty_prefix",
			words:    []string{"one", "two", "three"},
			word:     "",
			expected: []string{},
		},
		{
			title:    "single_letter_prefix_collision",
			words:    []string{"a", "ab", "abc", "abcd"},
			word:     "a",
			expected: []string{"a", "ab", "abc", "abcd"},
		},
		{
			title:    "prefix_is_substring_but_not_match",
			words:    []string{"lock", "keylock", "locker", "unlock"},
			word:     "lock",
			expected: []string{"lock", "locker"},
		},
		{
			title:    "unicode_characters",
			words:    []string{"🙂", "🙃", "😊", "😇smile", "😇happy"},
			word:     "😇",
			expected: []string{"😇smile", "😇happy"},
		},
		{
			title:    "japaneese",
			words:    []string{"ありがとう", "あります", "アニメ", "アイコン", "愛情", "愛"},
			word:     "あり",
			expected: []string{"ありがとう", "あります"},
		},
		{
			title:    "case_sensitive_check",
			words:    []string{"Apple", "app", "application", "APPlePie"},
			word:     "app",
			expected: []string{"app", "application"},
		},
		{
			title:    "word_equals_prefix_of_another",
			words:    []string{"run", "runner", "running", "runs"},
			word:     "run",
			expected: []string{"run", "runner", "running", "runs"},
		},
		{
			title:    "prefix_in_middle_should_not_match",
			words:    []string{"foobar", "barfoo", "foolish"},
			word:     "bar",
			expected: []string{"barfoo"},
		},
		{
			title:    "punctuation_and_symbols",
			words:    []string{"log-in", "log_out", "login!", "log", "log."},
			word:     "log",
			expected: []string{"log-in", "log_out", "login!", "log", "log."},
		},
		{
			title:    "space_separated_messages",
			words:    []string{"foobar barfoo foolish", "fuzz barz bazzing foo"},
			word:     "bar",
			expected: []string{"barfoo", "barz"},
		},
		{
			title:    "space_separated_query",
			words:    []string{"foobar barfoo foolish", "fuzz barz bazzing foo"},
			word:     "bar foo",
			expected: []string{"barfoo", "barz", "foobar", "foolish", "foo"},
		},
		{
			title:    "space_separated_query_intersecting_results",
			words:    []string{"foobar barfoo foolish", "fuzz barz bazzing foo", "found fork", "fone focus"},
			word:     "fo foo",
			expected: []string{"foobar", "foolish", "foo", "found", "fork", "fone", "focus"},
		},
		{
			title:    "word_partially_present",
			words:    []string{"cat", "cater", "cattle", "dog", "catalog"},
			word:     "catering",
			expected: []string{},
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			for _, word := range test.words {
				indexer.Index(word)
			}

			assert.ElementsMatch(t, test.expected, indexer.Search(test.word))
			indexer.Reset()
		})

	}

}
