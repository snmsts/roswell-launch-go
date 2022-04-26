package main

import (
	"os"
	"path"
	"strings"
)

func Launch(config map[string]string, args []string) {
	var program string
	_, file := path.Split(args[0])
	i := strings.Index(file, ".")
	if i != -1 {
		program = file[:i]
	} else {
		program = file
	}
	if program == "sbcl" {
		Launch_sbcl(config, args)
	} else {
		os.Exit(1)
	}
}
