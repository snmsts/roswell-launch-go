package arch

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

var unamem = ""

func UnameM() string {
	if unamem != "" {
		return unamem
	}
	str := ""

	if runtime.GOOS == "windows" {
		str = "x86-64"
	} else {
		out, err := exec.Command("uname", "-m").Output()
		if err != nil {
			log.Fatal(err)
		}
		str = string(out)
		str = strings.Trim(str, "\n")
	}
	if str == "i86pc" {
		str = "x86-64"
	} else if str == "i686" {
		str = "x86"
	} else if str == "i386" {
		str = "x86"
	} else if str == "amd64" {
		str = "x86-64"
	} else if str == "aarch64" {
		str = "arm64"
	} else if str == "armv6l" || str == "armv7l" {
		//char* result=system_("readelf -A /proc/self/exe |grep Tag_ABI_VFP_args|wc -l");
		//char* result2=remove_char("\r\n",result)
		//if (0==0) {
		str = "armhf"
		//}else {
		//str = "armel"
		//}
	} else if str == "armv5tejl" {
		str = "armel"
	} else {
		str = strings.Replace(str, "_", "-", -1)
	}

	unamem = str
	return unamem
}
