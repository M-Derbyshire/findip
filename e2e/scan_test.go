package e2e_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestProgramPrintsFoundIp(t *testing.T) {
	searchText := "test-search-text"
	expectedIp := "127.0.0.1"

	// Setup a server that we can cancel
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(searchText))
		}),
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logStr := fmt.Sprintf("error while running server as part of e2e testing: %v", err)
			fmt.Fprintln(os.Stderr, logStr)
		}
	}()
	defer server.Shutdown(context.Background())

	// Now run the program, looking for the endpoint we're running
	output, errOutput, exitStatus, err := runProgram("-from", expectedIp, "-to", expectedIp, "-port", "8080", "-searchtext", searchText)

	//assert
	if err != nil {
		t.Errorf("error while running command: %v", err)
		return
	}

	if len(errOutput) > 0 {
		t.Errorf("expected no output on stderr. got '%s'", errOutput)
	}

	if !strings.Contains(output, expectedIp) {
		t.Errorf("expected output to contain '%s'. got '%s'", expectedIp, output)
	}

	expectedExitCode := 0
	if exitStatus != expectedExitCode {
		t.Errorf("expected exit code to be %d. got %d", expectedExitCode, exitStatus)
	}
}
