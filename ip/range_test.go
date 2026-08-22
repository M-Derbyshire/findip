package ip_test

import (
	"findip/ip"
	util_test "findip/util"
	"fmt"
	"strings"
	"testing"
)

func TestGenerateIpAddressStringsToScanGeneratesTheExpectedIps(t *testing.T) {
	tests := []struct {
		name string
		from ip.IpAddress
		to   ip.IpAddress
	}{
		{
			name: "range1",
			from: ip.IpAddress{192, 168, 0, 0},
			to:   ip.IpAddress{195, 172, 5, 255},
		},
		{
			name: "range2",
			from: ip.IpAddress{254, 254, 254, 254},
			to:   ip.IpAddress{255, 255, 255, 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathToGolden := fmt.Sprintf("../snapshots/ip/generatesExpectedIps_%s.golden.txt", tt.name)

			results := ip.GenerateIpAddressStringsToScan(tt.from, tt.to)
			resultsSingle := strings.Join(results, "\n")
			resultsBytes := []byte(resultsSingle)

			util_test.ValidateSnapshot(t, resultsBytes, pathToGolden)
		})
	}
}
