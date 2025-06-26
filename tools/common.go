package tools

func CompareInt64(a, b int64) Comparison {
	switch {
	case a > b:
		return GREATER
	case a < b:
		return LESS
	default:
		return EQ
	}
}
