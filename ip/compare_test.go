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
