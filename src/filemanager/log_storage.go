package filemanager

import (
	"slices"

	"github.com/bodagovsky/logs_out/src/index"
	"github.com/bodagovsky/logs_out/tools"
)

// LogStorage represents single log file and related indexes
type LogStorage struct {
	tokenIndex *index.TokenIndex
	textIndex  *index.Indexer
	tsIndex    *index.TimestampIndex
}

func NewLogStorage(tokenIndex *index.TokenIndex, textIndex *index.Indexer, tsIndex *index.TimestampIndex) *LogStorage {
	return &LogStorage{
		tokenIndex: tokenIndex,
		textIndex:  textIndex,
		tsIndex:    tsIndex,
	}
}

func (ls LogStorage) Search(text string, from int64) []int64 {
	startOffset := ls.tsIndex.LocateLogEntry(from)

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
					result = append(result, off)
				}
			}
		}
	}
	if len(result) == 0 {
		result = append(result, startOffset)
	}

	slices.Sort(result)

	dedupResult := make([]int64, 0, len(result))

	for i := range result {
		if len(dedupResult) == 0 || result[i] > dedupResult[len(dedupResult)-1] {
			dedupResult = append(dedupResult, result[i])
		}
	}

	return dedupResult
}

func (ls *LogStorage) Index(message string, timestamp, offset int64) {
	ls.textIndex.Index(message)
	ls.tokenIndex.Index(message, offset)
	ls.tsIndex.InsertLogEntry(timestamp, offset)
}
