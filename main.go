package main

import (
	"os"
)

func main() {
	Launch(Parseconfig(ConfigPath()), os.Args)
}
