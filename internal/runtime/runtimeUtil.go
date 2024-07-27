package runtime

import (
	"os/exec"
	"syscall"
)

func IsExitCode(code int, err error) bool {
	if exitError, ok := err.(*exec.ExitError); ok {
		if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
			exitCode := status.ExitStatus()
			if exitCode == code {
				return true
			} else {
				return false
			}
		}
	} else {
		return false
	}
	return false
}
