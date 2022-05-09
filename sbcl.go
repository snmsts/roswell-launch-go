package main

import (
	"github.com/snmsts/roswell-launch-go/pkg/arch"
	"github.com/snmsts/roswell-launch-go/pkg/exec"
	"github.com/snmsts/roswell-launch-go/pkg/pwd"
	"os"
	"runtime"
)

func init() {
	Impl_map["sbcl"] = sbcl
}

func sbcl(config map[string]string) {
	args := os.Args
	bin_path := pwd.HomeDir()
	bin_path += "/" + "impls"
	bin_path += "/" + arch.CPU() // x86-64
	bin_path += "/" + arch.OS()  // darwin
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
