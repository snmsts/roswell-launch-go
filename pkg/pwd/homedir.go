package pwd

import (
	"os"
)

func HomeDir() string {
	rh := os.Getenv("ROSWELL_HOME")
	if rh != "" {
		return rh
	}
	return Systemhomedir() + "/.roswell"
}
