package index

type TokenIndex struct {
	tokens map[string][]int64
}

func NewTokenIndex() TokenIndex {
	return TokenIndex{
		tokens: make(map[string][]int64),
	}
}

func (ti *TokenIndex) InsertToken(token string, offset int64) {
	if _, ok := ti.tokens[token]; !ok {
		ti.tokens[token] = []int64{offset}
		return
	}
	ti.tokens[token] = append(ti.tokens[token], offset)
}
