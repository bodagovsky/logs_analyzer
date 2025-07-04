package filemanager

import (
	"github.com/bodagovsky/logs_out/src/index"
	"github.com/bodagovsky/logs_out/tools"
)

// LogStorage represents single log file and related indexes
type LogStorage struct {
	tokenIndex index.TokenIndex
	textIndex  index.Indexer
	tsIndex    index.TimestampIndex
}

func NewLogStorage(tokenIndex index.TokenIndex, textIndex index.Indexer, tsIndex index.TimestampIndex) *LogStorage {
	return &LogStorage{
		tokenIndex: tokenIndex,
		textIndex:  textIndex,
		tsIndex:    tsIndex,
	}
}

func (ls LogStorage) Search(text string, from int64) []int64 {
	startOffset := ls.tsIndex.LocateLogEntry(from)

	resultingOffsets := map[int64]struct{}{}
	var result []int64
	if len(text) > 0 {
		tokens := ls.textIndex.Search(text)
		if len(tokens) == 0 {
			// we did not find anything by text query, return empty response
			return result
		}
		for _, token := range tokens {
			offsetsByToken := ls.tokenIndex.GetOffsets(token)
			i := tools.Binary(offsetsByToken, startOffset, 0, len(offsetsByToken), tools.CompareInt64)
			for _, off := range offsetsByToken[i:] {
				if off >= startOffset {
					resultingOffsets[off] = struct{}{}
				}
			}
		}
	}
	if len(resultingOffsets) == 0 {
		result = append(result, startOffset)
	}

	for key := range resultingOffsets {
		result = append(result, key)
	}
	return result
}

func (ls *LogStorage) Index(message string, timestamp, offset int64) {
	ls.textIndex.Index(message)
	ls.tokenIndex.Index(message, offset)
	ls.tsIndex.InsertLogEntry(timestamp, offset)
}
