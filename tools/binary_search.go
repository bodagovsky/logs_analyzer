package tools

type Comparison int

const (
	LESS Comparison = iota
	EQ
	GREATER
)

// Optimized for searching in terms of low cardinality values
// returns first entry among multiple found targets
func Binary[T any](keys []T, target T, i, j int, compare func(T, T) Comparison) int {
	if i == len(keys) || compare(target, keys[j-1]) == GREATER {
		return j - 1
	}
	if j == 0 || compare(keys[i], target) == GREATER {
		return i
	}
	mid := (i + j) / 2
	if compare(keys[mid], target) == EQ {
		if mid > 0 && compare(keys[mid-1], target) == EQ {
			return Binary(keys, target, i, mid, compare)
		}
		return mid
	}
	if mid > 0 {
		if compare(target, keys[mid-1]) == GREATER && compare(target, keys[mid]) == LESS {
			return mid - 1
		}
	}
	if mid < len(keys)-1 {
		if compare(target, keys[mid]) == GREATER && compare(target, keys[mid+1]) == LESS {
			return mid
		}
	}
	if compare(target, keys[mid]) == GREATER {
		return Binary(keys, target, mid, j, compare)
	}
	return Binary(keys, target, i, mid, compare)
}
