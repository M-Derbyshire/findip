package config_test

import (
	"findip/config"
	"testing"
)

func TestCreateConfigWithDefaultsWillSetTheCorrectDefaults(t *testing.T) {
	expectedFromIp := []byte{192, 168, 0, 1}
	expectedToIp := []byte{192, 168, 0, 100}
	expectedPort := uint16(80)
	expectedEndpoint := ""
	expectedConcurrentRequests := 32
	expectedTextToFind := "hello world!"

	result := config.CreateConfigWithDefaults()

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
