package filemanager

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	types "github.com/bodagovsky/logs_out/src/types"
)


func LogsLookup(startOffset, from, to int64, files ...*os.File) ([]types.LogEntry, error) {
	var logs []types.LogEntry
	for _, file := range files {
		if startOffset >= 0 {
			file.Seek(startOffset, 0)

			reader := bufio.NewReader(file)
			line, err := reader.ReadString('\n')
			if err != nil {
				return logs, err
			}
			for len(line) > 0 {
				timestamp := strings.Trim(strings.Split(line, " ")[0], "[]")
				unixTs, err := strconv.Atoi(timestamp)
				if err != nil {
					return logs, err
				}				
				unix := time.Unix(int64(unixTs), 0).Unix()
				if unix >= from && unix < to {
					//TODO: add other field populating as well
					logs = append(logs, types.LogEntry{Timestamp: unix})
				}
				if unix >= to {
					break
				} 
				//TODO: handle error
				line, _ = reader.ReadString('\n')
			}
			startOffset = 0
		} else {
			break
		}
		
	}
	return logs, nil
}