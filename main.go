package main

import (
	"os"

	"github.com/bodagovsky/logs_out/src/datamanager"
)

func main() {
	file, err := os.Open("examples/data/client_42/logs/1750820538.log")
	if err != nil {
		panic(err)
	}
	stats, err := file.Stat()
	if err != nil {
		panic(err)
	}
	log := datamanager.LogEntry{}
	dm := datamanager.NewDataManager(file, stats.Size())

	dm.AppendLog(log)
	dm.Close()
}
