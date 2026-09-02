package e2e_test

import (
	"bytes"
	"os/exec"
)

func runProgram(commandArgs ...string) (outStr string, errStr string, exitCode int, err error) {
	cmd := exec.Command("./findip", commandArgs...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	runErr := cmd.Run()
	if runErr != nil {
		_, wasExitErr := runErr.(*exec.ExitError)

		if !wasExitErr {
			err = runErr
			return
		}
	}

	outStr = outb.String()
	errStr = errb.String()
	exitCode = cmd.ProcessState.ExitCode()
	return
}
