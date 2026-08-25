package ip_test

import (
	"findip/ip"
	"testing"
)

func TestIpAddressesAreEqualReturnsTrueForEquality(t *testing.T) {
	tests := []struct {
		name string
		ip   ip.IpAddress
	}{
		{
			name: "zeroes",
			ip:   ip.IpAddress{0, 0, 0, 0},
		},
		{
			name: "max",
			ip:   ip.IpAddress{0, 0, 0, 0},
		},
		{
			name: "local",
			ip:   ip.IpAddress{192, 168, 2, 241},
		},
		{
			name: "local",
			ip:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ip.IpAddressesAreEqual(tt.ip, tt.ip)

			if !result {
				t.Errorf("expected result to be true. got false")
			}
		})
	}
}

func TestIpAddressesAreEqualReturnsFalseIfOnlyOneIpIsNil(t *testing.T) {
	tests := []struct {
		name string
		ip1  ip.IpAddress
		ip2  ip.IpAddress
	}{
		{
			name: "ip1-nil",
			ip1:  nil,
			ip2:  ip.IpAddress{192, 168, 0, 1},
		},
		{
			name: "ip1-nil",
			ip1:  ip.IpAddress{192, 168, 0, 1},
			ip2:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ip.IpAddressesAreEqual(tt.ip1, tt.ip2)

			if result {
				t.Errorf("expected result to be false. got true")
			}
		})
	}
}

func TestIpAddressesAreEqualReturnsFalseIfIpAddressesDifferentLengths(t *testing.T) {
	ip1 := ip.IpAddress{192, 168, 0, 1}
	ip2 := ip.IpAddress{192, 168, 1}

	result := ip.IpAddressesAreEqual(ip1, ip2)

	if result {
		t.Errorf("expected result to be false. got true")
	}
}

func TestIpAddressesAreEqualReturnsFalseIfSegmentsNotEqual(t *testing.T) {
	tests := []struct {
		name string
		ip1  ip.IpAddress
		ip2  ip.IpAddress
	}{
		{
			name: "segment-1",
			ip1:  ip.IpAddress{192, 168, 1, 2},
			ip2:  ip.IpAddress{100, 168, 1, 2},
		},
		{
			name: "segment-2",
			ip1:  ip.IpAddress{192, 168, 1, 2},
			ip2:  ip.IpAddress{192, 100, 1, 2},
		},
		{
			name: "segment-3",
			ip1:  ip.IpAddress{192, 168, 1, 2},
			ip2:  ip.IpAddress{192, 168, 100, 2},
		},
		{
			name: "segment-4",
			ip1:  ip.IpAddress{192, 168, 1, 2},
			ip2:  ip.IpAddress{192, 168, 1, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ip.IpAddressesAreEqual(tt.ip1, tt.ip2)

			if result {
				t.Errorf("expected result to be false. got true")
			}
		})
	}
}

func TestIpAddressIsLessThanReturnsTrueIfLessThan(t *testing.T) {
	tests := []struct {
		name     string
		lowerIp  ip.IpAddress
		higherIp ip.IpAddress
	}{
		{
			name:     "segment-1",
			lowerIp:  ip.IpAddress{191, 168, 5, 10},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
		{
			name:     "segment-2",
			lowerIp:  ip.IpAddress{192, 167, 5, 10},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
		{
			name:     "segment-3",
			lowerIp:  ip.IpAddress{192, 168, 4, 10},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
		{
			name:     "segment-4",
			lowerIp:  ip.IpAddress{192, 168, 5, 9},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ip.IpAddressIsLessThan(tt.lowerIp, tt.higherIp)

			if err != nil {
				t.Errorf("expected error to be nil. got %v", err)
				return
			}

			if !result {
				t.Errorf("expected error to be true. got false")
			}
		})
	}
}

func TestIpAddressIsLessThanReturnsFalseIfEqual(t *testing.T) {
	ipAddr := ip.IpAddress{192, 168, 5, 10}

	result, err := ip.IpAddressIsLessThan(ipAddr, ipAddr)

	if err != nil {
		t.Errorf("expected error to be nil. got %v", err)
		return
	}

	if result {
		t.Errorf("expected result to be false. got true")
	}
}

func TestIpAddressIsLessThanReturnsFalseIfGreaterThan(t *testing.T) {
	tests := []struct {
		name     string
		lowerIp  ip.IpAddress
		higherIp ip.IpAddress
	}{
		{
			name:     "segment-1",
			lowerIp:  ip.IpAddress{191, 168, 5, 10},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
		{
			name:     "segment-2",
			lowerIp:  ip.IpAddress{192, 167, 5, 10},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
		{
			name:     "segment-3",
			lowerIp:  ip.IpAddress{192, 168, 4, 10},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
		{
			name:     "segment-4",
			lowerIp:  ip.IpAddress{192, 168, 5, 9},
			higherIp: ip.IpAddress{192, 168, 5, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// parameters passed in the wrong way round
			result, err := ip.IpAddressIsLessThan(tt.higherIp, tt.lowerIp)

			if err != nil {
				t.Errorf("expected error to be nil. got %v", err)
				return
			}

			if result {
				t.Errorf("expected result to be false. got true")
			}
		})
	}
}

func TestIpAddressIsLessThanReturnsErrorIfIpIsNil(t *testing.T) {
	expectedErr := "ip addresses must not be nil"

	tests := []struct {
		name     string
		lowerIp  ip.IpAddress
		higherIp ip.IpAddress
	}{
		{
			name:     "lower-ip-nil",
			lowerIp:  nil,
			higherIp: ip.IpAddress{192, 168, 0, 1},
		},
		{
			name:     "higher-ip-nil",
			lowerIp:  ip.IpAddress{192, 168, 0, 1},
			higherIp: nil,
		},
		{
			name:     "both-ip-nil",
			lowerIp:  nil,
			higherIp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ip.IpAddressIsLessThan(tt.lowerIp, tt.higherIp)

			if err == nil {
				t.Errorf("expected error to be '%s'. got nil", expectedErr)
				return
			}

			if err.Error() != expectedErr {
				t.Errorf("expected error to be '%s'. got '%v'", expectedErr, err)
			}
		})
	}
}

func TestIpAddressIsLessThanReturnsErrorIfSegmentLengthsDiffer(t *testing.T) {
	tests := []struct {
		name        string
		lowerIp     ip.IpAddress
		higherIp    ip.IpAddress
		expectedErr string
	}{
		{
			name:        "lower-ip-shorter",
			lowerIp:     ip.IpAddress{192, 168, 1},
			higherIp:    ip.IpAddress{192, 168, 0, 1},
			expectedErr: "ip addresses 192.168.1 and 192.168.0.1 do not have the same amount of segments",
		},
		{
			name:        "higher-ip-shorter",
			lowerIp:     ip.IpAddress{192, 168, 0, 1},
			higherIp:    ip.IpAddress{192, 168, 1},
			expectedErr: "ip addresses 192.168.0.1 and 192.168.1 do not have the same amount of segments",
		},
	}

	for _, tt := range tests {
		_, err := ip.IpAddressIsLessThan(tt.lowerIp, tt.higherIp)

		if err == nil {
			t.Errorf("expected error to be '%s'. got nil", tt.expectedErr)
			return
		}

		if err.Error() != tt.expectedErr {
			t.Errorf("expected error to be '%s'. got '%v'", tt.expectedErr, err)
		}
	}
}
