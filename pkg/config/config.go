package config

import (
	"encoding/csv"
	"github.com/snmsts/roswell-launch-go/pkg/pwd"
	"io"
	"os"
)

func Parse(fileName string) map[string]string {
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

func Path() string {
	path := pwd.HomeDir()
	path = path + "/config"
	return path
}
