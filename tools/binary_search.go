package tools

func Binary(keys []int64, target int64, i, j int) int {
	if i == len(keys) || target > keys[j-1] {
		return j - 1
	}
	if j == 0 || target < keys[i] {
		return i
	}
	mid := (i + j) / 2
	if keys[mid] == target {
		if mid > 0 && keys[mid-1] == target {
			return Binary(keys, target, i, mid)
		}
		return mid
	}
	if mid > 0 {
		if target > keys[mid-1] && target < keys[mid] {

			return mid - 1
		}
	}
	if mid < len(keys)-1 {
		if target > keys[mid] && target < keys[mid+1] {

			return mid
		}
	}
	if target > keys[mid] {
		return Binary(keys, target, mid, j)
	}
	return Binary(keys, target, i, mid)
}
