package config_test

import (
	"errors"
	"findip/config"
	"findip/ip"
	util_test "findip/util"
	"flag"
	"os"
	"testing"
)

func TestParseConfigFromFlagsWillSetTheCorrectDefaults(t *testing.T) {
	expectedFromIp := ip.IpAddress{192, 168, 0, 1}
	expectedToIp := ip.IpAddress{192, 168, 0, 100}
	expectedPort := uint(80)
	expectedEndpoint := ""
	expectedConcurrentRequests := uint(32)
	expectedTextToFind := "hello world!"

	result, err := config.ParseConfigFromFlags()
	if err != nil {
		t.Errorf("expected error to be nil. got %v", err)
		return
	}

	util_test.TestIpAddress(t, result.FromIp, expectedFromIp, "FromIp")
	util_test.TestIpAddress(t, result.ToIp, expectedToIp, "ToIp")

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
	expectedFromIp := ip.IpAddress{191, 167, 5, 10}
	expectedToIp := ip.IpAddress{193, 169, 6, 12}
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
	util_test.TestIpAddress(t, result.FromIp, expectedFromIp, "FromIp")
	util_test.TestIpAddress(t, result.ToIp, expectedToIp, "ToIp")

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
	tests := []struct {
		name string
		flag string
	}{
		{
			name: "from-ip",
			flag: "-from",
		},
		{
			name: "to-ip",
			flag: "-to",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			oldCommandLine := flag.CommandLine
			defer func() {
				os.Args = oldArgs
				flag.CommandLine = oldCommandLine
			}()

			os.Args = []string{
				"appname",
				tt.flag, "256.0.0.1",
			}
			flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)

			// run
			_, err := config.ParseConfigFromFlags()

			// assert
			if err == nil {
				t.Errorf("expected to get error. got nil")
				return
			}

			expectedErr := "ip address segment '256' in ip address '256.0.0.1' is not valid"
			if err.Error() != expectedErr {
				t.Errorf("expected error to be '%s'. got '%v'", expectedErr, err)
			}
		})
	}
}

func TestParseConfigFromFlagsWillReturnErrorIfFromIpGreaterThanToIp(t *testing.T) {
	// setup mocks
	oldArgs := os.Args
	oldCommandLine := flag.CommandLine
	originalCheck := ip.IpAddressIsLessThan
	defer func() {
		ip.IpAddressIsLessThan = originalCheck
		os.Args = oldArgs
		flag.CommandLine = oldCommandLine
	}()

	os.Args = []string{
		"appname",
	}
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)

	ip.IpAddressIsLessThan = func(lowerIp ip.IpAddress, higherIp ip.IpAddress) (bool, error) {
		return false, nil
	}

	// run
	_, err := config.ParseConfigFromFlags()

	// assert
	if err == nil {
		t.Error("expected error not to be nil")
		return
	}

	expectedErrorMessage := "-from ip address is greater than -to ip address"

	if err.Error() != expectedErrorMessage {
		t.Errorf("expected error to be '%s'. got '%v'", expectedErrorMessage, err)
	}
}

func TestParseConfigFromFlagsWillReturnErrorIfErrorWhenComparingIps(t *testing.T) {
	// setup mocks
	oldArgs := os.Args
	oldCommandLine := flag.CommandLine
	originalCheck := ip.IpAddressIsLessThan
	defer func() {
		ip.IpAddressIsLessThan = originalCheck
		os.Args = oldArgs
		flag.CommandLine = oldCommandLine
	}()

	os.Args = []string{
		"appname",
	}
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)

	expectedErrorMessage := "test error message 123"

	ip.IpAddressIsLessThan = func(lowerIp ip.IpAddress, higherIp ip.IpAddress) (bool, error) {
		return false, errors.New(expectedErrorMessage)
	}

	// run
	_, err := config.ParseConfigFromFlags()

	// assert
	if err == nil {
		t.Error("expected error not to be nil")
		return
	}

	if err.Error() != expectedErrorMessage {
		t.Errorf("expected error to be '%s'. got '%v'", expectedErrorMessage, err)
	}
}
