package index

import "strings"

/* Indexer is responsible for building full-text search index for logs */

type Indexer struct {
	head map[rune]*node
}

func NewIndexer() Indexer {
	return Indexer{head: make(map[rune]*node)}
}

func (I *Indexer) Index(message string) {
	var word strings.Builder

	for _, char := range message {
		if char == ' ' {
			if word.Len() > 0 {
				insert(word.String(), I.head)
				word.Reset()
			}
			continue
		}
		word.WriteRune(char)
	}
	if word.Len() > 0 {
		insert(word.String(), I.head)
	}
}

func (I *Indexer) Search(text string) []string {
	var word strings.Builder
	resultMap := make(map[string]struct{})

	for _, char := range text {
		if char == ' ' {
			if word.Len() > 0 {
				for _, result := range query(word.String(), I.head) {
					resultMap[result] = struct{}{}
				}
				word.Reset()
			}
			continue
		}
		word.WriteRune(char)
	}
	if word.Len() > 0 {
		for _, result := range query(word.String(), I.head) {
			resultMap[result] = struct{}{}
		}
	}
	result := make([]string, 0, len(resultMap))
	for key := range resultMap {
		result = append(result, key)
	}
	return result
}

func (I *Indexer) Reset() {
	I.head = make(map[rune]*node)
}

type node struct {
	value    rune
	entry    bool
	children map[rune]*node
}

func insert(word string, to map[rune]*node) {

	for i, r := range word {
		if _, ok := to[r]; !ok {
			to[r] = &node{value: r}
		}

		if to[r].children == nil {
			to[r].children = make(map[rune]*node)
		}
		to[r].entry = to[r].entry || len(string(r))+i == len(word)
		to = to[r].children

	}
}

func query(word string, from map[rune]*node) []string {
	if len(word) == 0 {
		return []string{}
	}

	for _, char := range word {
		if _, ok := from[char]; !ok {
			return []string{}
		}
		break
	}

	var entry bool
	for _, char := range word {
		if from[char].children == nil {
			break
		}
		entry = from[char].entry
		from = from[char].children
	}

	var output []string
	if entry {
		output = append(output, word)
	}
	for _, postfix := range gather(from) {
		output = append(output, word+postfix)
	}
	return output
}

func gather(children map[rune]*node) []string {
	var result []string
	for k, v := range children {
		if v.entry {
			result = append(result, string(k))
		}
		for _, word := range gather(v.children) {
			result = append(result, string(k)+word)
		}

	}
	return result
}
