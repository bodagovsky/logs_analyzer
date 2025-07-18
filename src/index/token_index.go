package index

import (
	"slices"
	"strings"
)

type TokenIndex struct {
	tokens map[string][]int64
}

func NewTokenIndex() TokenIndex {
	return TokenIndex{
		tokens: make(map[string][]int64),
	}
}

func (ti *TokenIndex) Index(message string, offset int64) {
	var token strings.Builder

	for _, char := range message {
		if char == ' ' {
			if token.Len() > 0 {
				tokenString := token.String()
				token.Reset()
				if _, ok := ti.tokens[tokenString]; !ok {
					ti.tokens[tokenString] = []int64{offset}
					continue
				}
				ti.tokens[tokenString] = append(ti.tokens[tokenString], offset)
			}
			continue
		}
		token.WriteRune(char)
	}
	if token.Len() > 0 {
		tokenString := token.String()
		if _, ok := ti.tokens[tokenString]; !ok {
			ti.tokens[tokenString] = []int64{offset}
			return
		}
		ti.tokens[tokenString] = append(ti.tokens[tokenString], offset)
	}
}

func (ti TokenIndex) GetOffsets(text string) []int64 {
	var token strings.Builder
	var result []int64

	for _, char := range text {
		if char == ' ' {
			if token.Len() > 0 {
				if offsets, ok := ti.tokens[token.String()]; ok {
					result = append(result, offsets...)
				}
				token.Reset()
			}
			continue
		}
		token.WriteRune(char)
	}

	if token.Len() > 0 {
		if offsets, ok := ti.tokens[token.String()]; ok {
			result = append(result, offsets...)
		}
	}

	// We need sorted results as later binary search is performed
	slices.Sort(result)
	return result
}

func (ti *TokenIndex) Reset() {
	ti.tokens = make(map[string][]int64)
}
