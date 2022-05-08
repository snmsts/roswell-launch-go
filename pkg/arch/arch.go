package arch

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

var unames = ""

func UnameS() string {
	if unames != "" {
		return unames
	}
	str := ""

	if runtime.GOOS == "windows" {
		str = "windows"
	} else {
		out, err := exec.Command("uname", "-s").Output()
		if err != nil {
			log.Fatal(err)
		}
		str = string(out)
		str = strings.Trim(str, "\n")
	}

	if str == "SunOS" {
		str = "solaris"
	} else {
		str = strings.ToLower(str)
	}
	unames = str
	return unames
}
