package filemanager

import (
	"bufio"
	"os"
)

// StreamManager performs operations on writing and retrieveng data from current stream
type StreamManager struct {
	// start holds the value of the first timestamp
	// so that StreamManager can decide when to rotate
	start int64

	// stream is the current log file accepting writing requests
	stream *os.File

	offset int64
}

func New(stream *os.File) *StreamManager {
	stat, err := stream.Stat()
	if err != nil {
		panic(err)
	}
	return &StreamManager{
		stream: stream,
		offset: stat.Size(),
	}
}

// AppendLog appends line to current stream and returns offset
func (sm *StreamManager) AppendLog(logEntry string) int64 {
	//TODO: add rotation logic
	offset := sm.offset
	if sm.offset > 0 {
		sm.stream.WriteString("\n")
		offset++
		sm.offset++
	}
	//TODO: add error work out
	n, _ := sm.stream.WriteString(logEntry)
	sm.offset += int64(n)
	return offset
}

// GetLinesByOffsets gets offsets and returns lines of stream by those
func (sm *StreamManager) GetLinesByOffsets(offsets []int64) []string {
	var lines []string
	var buf = bufio.NewReader(sm.stream)
	for _, offset := range offsets {
		sm.stream.Seek(offset, 0)
		line, _, _ := buf.ReadLine()
		lines = append(lines, string(line))
		buf.Reset(sm.stream)
	}
	sm.stream.Seek(sm.offset, 0)
	return lines
}
