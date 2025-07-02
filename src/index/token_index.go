package index

type TokenIndex struct {
	tokens map[string][]int64
}

func NewTokenIndex() TokenIndex {
	return TokenIndex{
		tokens: make(map[string][]int64),
	}
}

func (ti *TokenIndex) InsertToken(token string, line int64) {
	if _, ok := ti.tokens[token]; !ok {
		ti.tokens[token] = []int64{line}
		return
	}
	ti.tokens[token] = append(ti.tokens[token], line)
}

func (ti TokenIndex) GetOffsets (token string) []int64 {
	if offsets, ok := ti.tokens[token]; ok {
		return offsets
	}
	return []int64{}
}
