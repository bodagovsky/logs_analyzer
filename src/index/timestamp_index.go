package index

import "github.com/bodagovsky/logs_out/tools"

// TIMESPAN specifies window in which timetsamps are aggregated in index
const TIMESPAN int64 = 3600 // 1 hour

type TimestampIndex struct {
	keys     []int64
	mappings map[int64]int64
}

func NewTimestampIndex() TimestampIndex {
	return TimestampIndex{
		keys:     []int64{},
		mappings: make(map[int64]int64),
	}
}

func (ti *TimestampIndex) InsertLogEntry(ts, line int64) {
	if len(ti.keys) == 0 {
		ti.keys = append(ti.keys, ts)
		ti.mappings[ts] = line
		return
	}
	i := len(ti.keys) - 1
	if ts-ti.keys[i] >= TIMESPAN {
		ti.keys = append(ti.keys, ts)
		ti.mappings[ts] = line
	}
}

func (ti TimestampIndex) LocateLogEntry(from int64) int64 {
	i := tools.Binary(ti.keys, from, 0, len(ti.keys), tools.CompareInt64)
	return ti.mappings[ti.keys[i]]
}

func (ti *TimestampIndex) Reset() {
	ti.keys = []int64{}
	ti.mappings = make(map[int64]int64)
}
