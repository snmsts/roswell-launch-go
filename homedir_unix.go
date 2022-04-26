package main

import (
	"github.com/nogproject/nog/backend/pkg/pwd"
	"os"
)

func Systemhomedir() string {
	user := os.Getenv("SUDO_USER")
	uid := os.Getuid()
	var pw *pwd.Passwd

	if user != "" && uid == 0 {
		pw = pwd.Getpwnam(user)
	} else {
		pw = pwd.Getpwuid(uint32(uid))
	}
	if pw == nil {
		return ""
	}
	//fmt.Print("homedir is'",pw.Dir,"' and SUDO user is'",user,"'\n")
	return pw.Dir
}

func HomeDir() string {
	rh := os.Getenv("ROSWELL_HOME")
	if rh != "" {
		return rh
	}
	return Systemhomedir() + "/.roswell"
}

func ConfigPath() string {
	path := HomeDir()
	path = path + "/config"
	return path
}
