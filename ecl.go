package main

import (
	"github.com/snmsts/roswell-launch-go/pkg/arch"
	"github.com/snmsts/roswell-launch-go/pkg/exec"
	"github.com/snmsts/roswell-launch-go/pkg/pwd"
	"os"
	"runtime"
)

func init() {
	Impl_map["ecl"] = ecl
}

func ecl(config map[string]string) {
	args := os.Args
	bin_path := pwd.HomeDir()
	bin_path += "/" + "impls"
	bin_path += "/" + arch.CPU() // x86-64
	bin_path += "/" + arch.OS()  // darwin
	bin_path += "/" + "ecl"
	bin_path += "/" + config["ecl.version"]
	bin_path += "/" + "bin"
	bin_path += "/" + "ecl"
	if runtime.GOOS == "windows" {
		bin_path += ".exe"
	}
	args[0] = bin_path
	exec.Exec(bin_path, args, nil)
}
