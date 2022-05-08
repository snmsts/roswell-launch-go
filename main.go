package main

import (
	"github.com/snmsts/roswell-launch-go/pkg/config"
	"os"
	"path"
	"strings"
)

var Impl_map = make(map[string]func(config map[string]string))

func remove_extention(full_path string) string {
	_, file := path.Split(full_path)
	i := strings.Index(file, ".")
	if i != -1 {
		return file[:i]
	} else {
		return file
	}
}

func main() {
	program := remove_extention(os.Args[0])

	if Impl_map[program] != nil {
		Impl_map[program](config.Parse(config.Path()))
	} else {
		os.Exit(1)
	}
}
