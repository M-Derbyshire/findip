package e2e_test

import (
	"strings"
	"testing"
)

func TestProgramPrintsInvalidConfigErrorWithCorrectStatusCode(t *testing.T) {
	output, errOutput, exitStatus, err := runProgram("-from", "192.168.0.2", "-to", "192.168.0.1")
	if err != nil {
		t.Errorf("error while running command: %v", err.Error())
		return
	}

	if len(output) > 0 {
		t.Errorf("expected no output on stdout. got '%s'", output)
	}

	expectedErrText := "-from ip address is greater than -to ip address"

	if !strings.Contains(errOutput, expectedErrText) {
		t.Errorf("expected error output to contain '%v'. got '%v'", expectedErrText, errOutput)
	}

	expectedExitCode := 1

	if exitStatus != expectedExitCode {
		t.Errorf("expected exit code to be %d. got %d", expectedExitCode, exitStatus)
	}
}

func TestProgramPrintsIpNotFoundErrorWithCorrectStatusCode(t *testing.T) {
	output, errOutput, exitStatus, err := runProgram("-from", "0.0.0.0", "-to", "0.0.0.0")
	if err != nil {
		t.Errorf("error while running command: %v", err.Error())
		return
	}

	if len(output) > 0 {
		t.Errorf("expected no output on stdout. got '%s'", output)
	}

	expectedErrText := "unable to find ip"

	if !strings.Contains(errOutput, expectedErrText) {
		t.Errorf("expected error output to contain '%v'. got '%v'", expectedErrText, errOutput)
	}

	expectedExitCode := 0

	if exitStatus != expectedExitCode {
		t.Errorf("expected exit code to be %d. got %d", expectedExitCode, exitStatus)
	}
}
