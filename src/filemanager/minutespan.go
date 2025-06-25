package filemanager

// minuteSpan stores mappings for timestamps by each minute within one hour
type MinuteSpan struct {
	// keys stores all the entries from mapping in non-decreasing order
	keys []int64

	// mapping stores start timestamps by each minute as key and an offset
	// for that period where the offset is the line number
	mapping map[int64]int64
}

func (s *MinuteSpan) InsertLogEntry(ts, line int64) {
	if len(s.keys) == 0 {
		s.keys = append(s.keys, ts)
		s.mapping[ts] = line
		return
	}
	i := len(s.keys) - 1
	if ts-s.keys[i] > 60 {
		s.keys = append(s.keys, ts)
		s.mapping[ts] = line
	}
}

func (s MinuteSpan) LocateLogEntries(from int64) int64 {
	// step 1: perform binary search to find the line where desired timeline starts
	i := Binary(s.keys, from, 0, len(s.keys), false)
	if i == -1 {
		return -1
	}

	// step 2: extract the offset
	return s.mapping[s.keys[i]]
}

// Binary seacrhes sorted keys in binary seacrh manner
// in case of multiple targets present in keys return the index of a first occurence
func Binary(keys []int64, target int64, i, j int, upto bool) int {
	if j == 0 || target > keys[j-1] {
		if upto {
			return j - 1
		}
		return -1
	}
	if i == len(keys) || target < keys[i] {
		if upto {
			return -1
		}
		return i
	}
	mid := (i + j) / 2
	if keys[mid] == target {
		if mid > 0 && keys[mid-1] == target {
			return Binary(keys, target, i, mid, upto)
		}
		return mid
	}
	if mid > 0 {
		if target > keys[mid-1] && target < keys[mid] {
			if upto {
				return mid - 1
			}
			return mid
		}
	}
	if mid < len(keys)-1 {
		if target > keys[mid] && target < keys[mid+1] {
			if upto {
				return mid
			}
			return mid + 1
		}
	}
	if target > keys[mid] {
		return Binary(keys, target, mid, j, upto)
	}
	return Binary(keys, target, i, mid, upto)
}
