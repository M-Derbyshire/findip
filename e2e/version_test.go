package e2e_test

import (
	"findip/info"
	"fmt"
	"testing"
)

func TestProgramPrintsVersiontext(t *testing.T) {
	tests := []string{
		"-v",
		"-version",
	}

	for _, flagText := range tests {
		t.Run(flagText, func(t *testing.T) {
			output, errOutput, exitStatus, err := runProgram(flagText)

			if err != nil {
				t.Errorf("error while running command: %v", err)
				return
			}

			if len(errOutput) > 0 {
				t.Errorf("expected no output on stderr. got '%s'", errOutput)
			}

			expectedOutput := fmt.Sprintf("%s\n", info.Version)
			if output != expectedOutput {
				t.Errorf("expected output to be '%s'. got '%s'", expectedOutput, output)
			}

			expectedExitCode := 0
			if exitStatus != expectedExitCode {
				t.Errorf("expected exit code to be %d. got %d", expectedExitCode, exitStatus)
			}
		})
	}
}
