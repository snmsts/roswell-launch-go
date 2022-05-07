package main

import (
	"github.com/snmsts/roswell-launch-go/pkg/exec"
	"runtime"
)

func Launch_sbcl(config map[string]string, args []string) {
	bin_path := HomeDir()
	bin_path += "/" + "impls"
	bin_path += "/" + UnameM() // x86-64
	bin_path += "/" + UnameS() // darwin
	bin_path += "/" + "sbcl-bin"
	bin_path += "/" + config["sbcl-bin.version"]
	bin_path += "/" + "bin"
	bin_path += "/" + "sbcl"
	if runtime.GOOS == "windows" {
		bin_path += ".exe"
	}
	args[0] = bin_path
	exec.Exec(bin_path, args, nil)
}
