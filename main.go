package main

import (
	"os"
	"path"
	"strings"
)

var Impl_map = make(map[string]func(config map[string]string, args []string))

func main() {
	config := Parseconfig(ConfigPath())
	args := os.Args
	var program string
	_, file := path.Split(args[0])
	i := strings.Index(file, ".")
	if i != -1 {
		program = file[:i]
	} else {
		program = file
	}
	if Impl_map[program] != nil {
		Impl_map[program](config, args)
	} else {
		os.Exit(1)
	}
}
