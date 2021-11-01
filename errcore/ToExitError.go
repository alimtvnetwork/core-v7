package errcore

import "os/exec"

func ToExitError(err error) *exec.ExitError {
	if err == nil {
		return nil
	}

	exitError, _ := err.(*exec.ExitError)

	return exitError
}
