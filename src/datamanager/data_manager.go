package datamanager

import (
	"os"
	"strings"

	"github.com/bodagovsky/logs_out/src/index"
)


type DataManager struct {
	tokenIndex index.TokenIndex
	textIndex index.TextIndex
	file *os.File
	offset int64
	filename string
}

func NewDataManager (file *os.File, offset int64) DataManager {
	return DataManager{
		file: file,
		offset: offset,
		filename: file.Name(),
	}
}

func (dm DataManager) Close () {
	dm.file.Close()
}

func (dm DataManager) QueryLogs (q Query) []string {

	return []string{}
}


func (dm *DataManager) AppendLog (logEntry LogEntry) error {	
	n, err := dm.file.Write([]byte(logEntry.Format()))
	if err != nil {return err}

	// TODO: get rid of split, index words by skipping delimiter
	words := strings.Split(logEntry.Text, " ")

	for i := range words {
		dm.textIndex.Index(words[i])
		dm.tokenIndex.InsertToken(words[i], dm.offset)
	}
	dm.offset += int64(n)
	return nil
}