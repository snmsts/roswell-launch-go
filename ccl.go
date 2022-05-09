package main

import (
	"github.com/snmsts/roswell-launch-go/pkg/arch"
	"github.com/snmsts/roswell-launch-go/pkg/exec"
	"github.com/snmsts/roswell-launch-go/pkg/pwd"
	"os"
	"runtime"
)

func init() {
	Impl_map["ccl"] = ccl
}

func ccl_binname() string {
	cpu := arch.CPU()
	os := arch.OS()
	var ret string

	if os == "linux" {
		if cpu != "armhf" {
			ret = "l"
		} else {
			ret = ""
		}
	} else if os == "windows" {
		ret = "w"
	} else if os == "darwin" {
		ret = "d"
	} else if os == "freebsd" {
		ret = "f"
	} else if os == "solaris" {
		ret = "s"
	}

	if cpu == "x86-64" || cpu == "x86" {
		ret = ret + "x86"
	} else if cpu == "armhf" {
		ret = ret + "arm"
	}
	ret = ret + "cl"
	if cpu != "x86" {
		ret = ret + "64"
	}
	return ret
}

func ccl(config map[string]string) {
	args := os.Args
	bin_path := pwd.HomeDir()
	bin_path += "/" + "impls"
	bin_path += "/" + arch.CPU()
	bin_path += "/" + arch.OS()
	bin_path += "/" + "ccl-bin"
	bin_path += "/" + config["ccl-bin.version"]
	bin_path += "/" + ccl_binname()
	if runtime.GOOS == "windows" {
		bin_path += ".exe"
	}
	args[0] = bin_path
	exec.Exec(bin_path, args, nil)
}
