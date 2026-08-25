package ip

import (
	"errors"
	"fmt"
)

func IpAddressesAreEqual(ip1 IpAddress, ip2 IpAddress) bool {
	ip1IsNil := ip1 == nil
	ip2IsNil := ip2 == nil

	if ip1IsNil != ip2IsNil {
		return false
	}

	if len(ip1) != len(ip2) {
		return false
	}

	for i, segment1 := range ip1 {
		segment2 := ip2[i]

		if segment1 != segment2 {
			return false
		}
	}

	return true
}

func IpAddressIsLessThan(lowerIp IpAddress, higherIp IpAddress) (bool, error) {
	if lowerIp == nil || higherIp == nil {
		return false, errors.New("ip addresses must not be nil")
	}

	if len(lowerIp) != len(higherIp) {
		lowerStr := convertIpAddressToString(lowerIp)
		higherStr := convertIpAddressToString(higherIp)
		return false, fmt.Errorf("ip addresses %s and %s do not have the same amount of segments", lowerStr, higherStr)
	}

	for i, lowerSegment := range lowerIp {
		higherSegment := higherIp[i]

		if lowerSegment == higherSegment {
			continue
		}

		if lowerSegment < higherSegment {
			return true, nil
		}

		return false, nil
	}

	// all segments must be equal
	return false, nil
}
