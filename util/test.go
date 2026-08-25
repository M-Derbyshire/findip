package util_test

import (
	"findip/ip"
	"flag"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var updateTestSnapshots = flag.Bool("update", false, "update golden snaphot files")

func ValidateSnapshot(t *testing.T, resultBytes []byte, snapshotPath string) {
	if *updateTestSnapshots {
		err := os.WriteFile(snapshotPath, resultBytes, 0644)
		if err != nil {
			t.Errorf("error while updating snapshot: '%v'", err)
			return
		}
	}

	expected, err := os.ReadFile(snapshotPath)
	if err != nil {
		t.Errorf("error while reading snapshot file: '%v'", err)
		return
	}

	diff := cmp.Diff(string(expected), string(resultBytes))
	if diff != "" {
		t.Errorf("result did not match the snapshot:\n%s", diff)
	}
}

func TestIpAddress(t *testing.T, resultIp ip.IpAddress, expectedIp ip.IpAddress, propertyName string) {
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
