//go:build windows
// +build windows

package main

import (
	"github.com/snmsts/roswell-launch-go/pkg/pwd"
	"log"
	"os"
)

func Systemhomedir() string {
	ret, err := pwd.ProfileFolder()
	if err != nil {
		log.Fatal(err)
	}
	return ret
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
