//go:build windows

package exec

import (
	"os"
	"os/exec"
)

func Exec(exec_path string, args []string, env []string) {
	binary, lookErr := exec.LookPath(exec_path)
	if lookErr != nil {
		panic(lookErr)
	}

	if env == nil {
		env = os.Environ()
	}

	cmd := exec.Command(binary)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Args = args
	cmd.Env = env

	cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()
	os.Exit(exitCode)
}
