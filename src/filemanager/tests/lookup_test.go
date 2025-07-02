package tests

import (
	"os"
	"testing"

	"github.com/bodagovsky/logs_out/src/filemanager"
	"github.com/stretchr/testify/assert"
)

func TestLookup_SingleFile(t *testing.T) {
	file, err := os.Open("data/lookup/1750820538.log")

	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	logs, err := filemanager.LogsLookup(0, 1750820538, 1750820540, file)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(logs))

	expected := int64(1750820538)
	for _, log := range logs {
		assert.Equal(t, expected, log.Timestamp)
		expected++
	}
}

func TestLookup_TwoFiles(t *testing.T) {
	file_01, err := os.Open("data/lookup/1750820538.log")

	if err != nil {
		t.Fatal(err)
	}
	defer file_01.Close()

	file_02, err := os.Open("data/lookup/1750906938.log")

	if err != nil {
		t.Fatal(err)
	}
	defer file_02.Close()

	logs, err := filemanager.LogsLookup(0, 1750820538, 1750906940, file_01, file_02)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 5, len(logs))

	expected := []int64{
		1750820538,
		1750820539,
		1750820540,
		1750906938,
		1750906939,
	}
	i := 0
	for _, log := range logs {
		assert.Equal(t, expected[i], log.Timestamp)
		i++
	}
}

func TestLookup_SingleFile_V2(t *testing.T) {
	file, err := os.Open("data/lookup/1750820538.log")

	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	logs, err := filemanager.LogsLookup(0, 1750820538, 1750820550, file)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(logs))

	expected := int64(1750820538)
	for _, log := range logs {
		assert.Equal(t, expected, log.Timestamp)
		expected++
	}
}

func TestLookup_TwoFiles_FirstFileNotRelevant(t *testing.T) {
	file_01, err := os.Open("data/lookup/1750820538.log")

	if err != nil {
		t.Fatal(err)
	}
	defer file_01.Close()

	file_02, err := os.Open("data/lookup/1750906938.log")

	if err != nil {
		t.Fatal(err)
	}
	defer file_02.Close()

	logs, err := filemanager.LogsLookup(0, 1750820541, 1750906940, file_01, file_02)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(logs))

	expected := int64(1750906938)
	for _, log := range logs {
		assert.Equal(t, expected, log.Timestamp)
		expected++
	}
}

func TestLookup_SingleFile_EmptyResult(t *testing.T) {
	file, err := os.Open("data/lookup/1750820538.log")

	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	logs, err := filemanager.LogsLookup(0, 1750820542, 1750820543, file)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, len(logs))
}
