package ip_test

import (
	"findip/ip"
	"testing"
)

func TestParseIpV4StringWillErrorIfInvalidIpAddress(t *testing.T) {
	tests := []struct {
		name              string
		ipStr             string
		expectedErrorText string
	}{
		{
			name:              "negative segment",
			ipStr:             "-100.0.0.0",
			expectedErrorText: "ip address segment '-100' in ip address '-100.0.0.0' is not valid",
		},
		{
			name:              "segment 0 too large",
			ipStr:             "256.0.0.0",
			expectedErrorText: "ip address segment '256' in ip address '256.0.0.0' is not valid",
		},
		{
			name:              "segment 1 too large",
			ipStr:             "1.256.0.0",
			expectedErrorText: "ip address segment '256' in ip address '1.256.0.0' is not valid",
		},
		{
			name:              "segment 2 too large",
			ipStr:             "1.0.256.0",
			expectedErrorText: "ip address segment '256' in ip address '1.0.256.0' is not valid",
		},
		{
			name:              "segment 3 too large",
			ipStr:             "1.0.0.256",
			expectedErrorText: "ip address segment '256' in ip address '1.0.0.256' is not valid",
		},
		{
			name:              "too many segments",
			ipStr:             "255.0.0.1.1",
			expectedErrorText: "ip address '255.0.0.1.1' should have 4 segments",
		},
		{
			name:              "too few segments",
			ipStr:             "255.0.1",
			expectedErrorText: "ip address '255.0.1' should have 4 segments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ip.ParseIpV4String(tt.ipStr)

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

func TestParseIpV4StringWillParseIpAddressesCorrectly(t *testing.T) {
	tests := []struct {
		name       string
		ipStr      string
		expectedIp ip.IpAddress
	}{
		{
			name:       "low",
			ipStr:      "1.0.0.1",
			expectedIp: ip.IpAddress{1, 0, 0, 1},
		},
		{
			name:       "mid",
			ipStr:      "192.168.0.10",
			expectedIp: ip.IpAddress{192, 168, 0, 10},
		},
		{
			name:       "high",
			ipStr:      "255.255.255.255",
			expectedIp: ip.IpAddress{255, 255, 255, 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ip.ParseIpV4String(tt.ipStr)

			if err != nil {
				t.Errorf("expected error to be nil. got %v", err)
				return
			}

			if len(result) != 4 {
				t.Errorf("expected result ip to have 4 segments. got %d", len(result))
				return
			}

			for i, segment := range result {
				expectedSegment := tt.expectedIp[i]

				if segment != expectedSegment {
					t.Errorf("expected ip segment %d to be %d. got %d", i, expectedSegment, segment)
				}
			}
		})
	}
}
