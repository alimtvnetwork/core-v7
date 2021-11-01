package iserror

import "os/exec"

func ExitError(err error) bool {
	_, isOkay := err.(*exec.ExitError)

	return isOkay
}
