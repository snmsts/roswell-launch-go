package main

import (
	"encoding/csv"
	"io"
	"os"
)

func Parseconfig(fileName string) map[string]string {
	result := make(map[string]string)
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comma = '\t'
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		result[record[0]] = record[2]
	}
	return result
}
