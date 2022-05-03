package main

import (
	"github.com/snmsts/roswell-launch-go/pkg/pwd"
	"os"
)

func HomeDir() string {
	rh := os.Getenv("ROSWELL_HOME")
	if rh != "" {
		return rh
	}
	return pwd.Systemhomedir() + "/.roswell"
}

func ConfigPath() string {
	path := HomeDir()
	path = path + "/config"
	return path
}