// build !windows

package pwd

import (
	"log"
	"os"
	"os/user"
	"strconv"
)

func Systemhomedir() string {
	sudo_user := os.Getenv("SUDO_USER")
	uid := os.Getuid()

	if sudo_user != "" && uid == 0 {
		userInfo, err := user.Lookup(sudo_user)
		if err != nil {
			log.Panicf("no such user", err)
		}
		return userInfo.HomeDir
	} else {
		uid_str := strconv.Itoa(uid)
		userInfo, err := user.LookupId(uid_str)
		if err != nil {
			log.Panicf("no such user", err)
		}
		return userInfo.HomeDir
	}
	log.Panic("no such user")
	return ""
}
