package config_test

import (
	"findip/config"
	"flag"
	"os"
	"testing"
)

func TestParseConfigFromFlagsWillSetTheCorrectDefaults(t *testing.T) {
	expectedFromIp := []byte{192, 168, 0, 1}
	expectedToIp := []byte{192, 168, 0, 100}
	expectedPort := uint(80)
	expectedEndpoint := ""
	expectedConcurrentRequests := uint(32)
	expectedTextToFind := "hello world!"

	result, err := config.ParseConfigFromFlags()
	if err != nil {
		t.Errorf("expected error to be nil. got %v", err)
		return
	}

	testIpAddress(t, result.FromIp, expectedFromIp, "FromIp")
	testIpAddress(t, result.ToIp, expectedToIp, "ToIp")

	if result.Port != expectedPort {
		t.Errorf("expected port to be %d. got %d", expectedPort, result.Port)
	}

	if result.Endpoint != expectedEndpoint {
		t.Errorf("expected endpoint to be '%s'. got '%s'", expectedEndpoint, result.Endpoint)
	}

	if result.ConcurrentRequests != expectedConcurrentRequests {
		t.Errorf("expected concurrent requests to be %d. got %d", expectedConcurrentRequests, result.ConcurrentRequests)
	}

	if result.TextToFind != expectedTextToFind {
		t.Errorf("expected text-to-find to be '%s'. got '%s'", expectedTextToFind, result.TextToFind)
	}
}

func TestParseConfigFromFlagsWillParseCommandLineFlagValues(t *testing.T) {
	expectedFromIp := []byte{191, 167, 5, 10}
	expectedToIp := []byte{193, 169, 6, 12}
	expectedPort := uint(81)
	expectedEndpoint := "/test"
	expectedConcurrentRequests := uint(64)
	expectedTextToFind := "hello test!"

	// setup mocks
	oldArgs := os.Args
	oldCommandLine := flag.CommandLine
	defer func() {
		os.Args = oldArgs
		flag.CommandLine = oldCommandLine
	}()

	os.Args = []string{
		"appname",
		"-from", "191.167.5.10",
		"-to", "193.169.6.12",
		"-port", "81",
		"-endpoint", expectedEndpoint,
		"-concurrentrequests", "64",
		"-searchtext", expectedTextToFind,
	}

	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)

	// run
	result, err := config.ParseConfigFromFlags()
	if err != nil {
		t.Errorf("expected error to be nil. got %v", err)
		return
	}

	// assert
	testIpAddress(t, result.FromIp, expectedFromIp, "FromIp")
	testIpAddress(t, result.ToIp, expectedToIp, "ToIp")

	if result.Port != expectedPort {
		t.Errorf("expected port to be %d. got %d", expectedPort, result.Port)
	}

	if result.Endpoint != expectedEndpoint {
		t.Errorf("expected endpoint to be '%s'. got '%s'", expectedEndpoint, result.Endpoint)
	}

	if result.ConcurrentRequests != expectedConcurrentRequests {
		t.Errorf("expected concurrent requests to be %d. got %d", expectedConcurrentRequests, result.ConcurrentRequests)
	}

	if result.TextToFind != expectedTextToFind {
		t.Errorf("expected text-to-find to be '%s'. got '%s'", expectedTextToFind, result.TextToFind)
	}
}

func TestParseConfigFromFlagsWillReturnErrorIfInvalidIpAddress(t *testing.T) {
	validIp := "192.168.0.1"

	tests := []struct {
		name              string
		fromIpStr         string
		toIpStr           string
		expectedErrorText string
	}{
		{
			name:              "negative segment",
			fromIpStr:         "-100.0.0.0",
			toIpStr:           validIp,
			expectedErrorText: "ip address segment '-100' in ip address '-100.0.0.0' is not valid",
		},
		{
			name:              "segment 0 too large",
			fromIpStr:         "256.0.0.0",
			toIpStr:           validIp,
			expectedErrorText: "ip address segment '256' in ip address '256.0.0.0' is not valid",
		},
		{
			name:              "segment 1 too large",
			fromIpStr:         "1.256.0.0",
			toIpStr:           validIp,
			expectedErrorText: "ip address segment '256' in ip address '1.256.0.0' is not valid",
		},
		{
			name:              "segment 2 too large",
			fromIpStr:         "1.0.256.0",
			toIpStr:           validIp,
			expectedErrorText: "ip address segment '256' in ip address '1.0.256.0' is not valid",
		},
		{
			name:              "segment 3 too large",
			fromIpStr:         "1.0.0.256",
			toIpStr:           validIp,
			expectedErrorText: "ip address segment '256' in ip address '1.0.0.256' is not valid",
		},
		{
			name:              "segment 0 in to-ip too large",
			fromIpStr:         validIp,
			toIpStr:           "256.0.0.1",
			expectedErrorText: "ip address segment '256' in ip address '256.0.0.1' is not valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup mocks
			oldArgs := os.Args
			oldCommandLine := flag.CommandLine
			defer func() {
				os.Args = oldArgs
				flag.CommandLine = oldCommandLine
			}()

			os.Args = []string{
				"appname",
				"-from", tt.fromIpStr,
				"-to", tt.toIpStr,
			}

			flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)

			// run
			_, err := config.ParseConfigFromFlags()

			// assert
			if err == nil {
				t.Errorf("expected error to '%s'. got nil", tt.expectedErrorText)
				return
			}

			if err.Error() != tt.expectedErrorText {
				t.Errorf("expected error to be '%s'. got '%v'", tt.expectedErrorText, err)
			}
		})
	}
}

func testIpAddress(t *testing.T, resultIp []byte, expectedIp []byte, propertyName string) {
	if len(resultIp) != len(expectedIp) {
		t.Errorf("expected %s length to be %d. got %d", propertyName, len(expectedIp), len(resultIp))
	} else {
		for i, val := range resultIp {
			expected := expectedIp[i]
			if val != expected {
				t.Errorf("expected IP block at index %d for %s to be %d. got %d", i, propertyName, expected, val)
			}
		}
	}
}
