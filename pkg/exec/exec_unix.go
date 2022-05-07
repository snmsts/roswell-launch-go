//go:build !windows

package exec

import (
	"os"
	"os/exec"
	"syscall"
)

func Exec(exec_path string, args []string, env []string) {
	binary, lookErr := exec.LookPath(exec_path)
	if lookErr != nil {

		panic(lookErr)
	}
	if env == nil {
		env = os.Environ()
	}
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
