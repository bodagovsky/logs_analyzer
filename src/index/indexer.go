package index

/* Indexer is responsible for building full-text search index for logs */

type Indexer struct {
	head map[byte]*node
}

func NewIndexer() *Indexer {
	return &Indexer{head: make(map[byte]*node)}
}

func (I *Indexer) Index(word string) {
	insert(word, I.head)
}

func (I *Indexer) Search(word string) []string {
	return query(word, I.head)
}

func (I *Indexer) Reset() {
	I.head = make(map[byte]*node)
}

type node struct {
	value    byte
	entry    bool
	children map[byte]*node
}

func insert(word string, to map[byte]*node) {
	for i := range word {
		if _, ok := to[word[i]]; !ok {
			to[word[i]] = &node{value: word[i]}
		}

		to[word[i]].entry = i == len(word)-1 || to[word[i]].entry

		if to[word[i]].children == nil {
			to[word[i]].children = make(map[byte]*node)
		}
		to = to[word[i]].children
	}
}

func query(word string, from map[byte]*node) []string {
	if len(word) == 0 {
		return []string{}
	}
	if _, ok := from[word[0]]; !ok {
		return []string{}
	}
	var entry bool
	for i := range word {
		if from[word[i]].children == nil {
			break
		}
		entry = from[word[i]].entry
		from = from[word[i]].children
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

func gather(children map[byte]*node) []string {
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
