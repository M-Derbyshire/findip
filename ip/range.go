package ip

import (
	"fmt"
	"strings"
)

var GenerateIpAddressStrings = func(fromIp IpAddress, toIp IpAddress) []string {
	currentIp := make(IpAddress, len(fromIp))
	copy(currentIp, fromIp)

	ipsToScan := make([]string, 0, 1024) // capacity here is just a high value, to save on memory reshuffling in most cases

	lastIpHasBeenProcessed := false
	for !lastIpHasBeenProcessed {
		incrementIpAddress(&currentIp)

		ipStr := convertIpAddressToString(currentIp)
		ipsToScan = append(ipsToScan, ipStr)

		lastIpHasBeenProcessed = ipAddressesAreEqual(currentIp, toIp)
	}

	return ipsToScan
}

func incrementIpAddress(ip *IpAddress) {
	// the segment index to increment (e.g. a value of 3 will edit the 4th segment -- the 10 in 192.168.0.10)
	segmentToIncrement := len(*ip) - 1

	for {
		if (*ip)[segmentToIncrement] == 255 {
			(*ip)[segmentToIncrement] = 0
			segmentToIncrement--
		} else {
			(*ip)[segmentToIncrement]++
			return
		}
	}
}

func ipAddressesAreEqual(ip1 IpAddress, ip2 IpAddress) bool {
	ip1IsNil := ip1 == nil
	ip2IsNil := ip2 == nil

	if ip1IsNil != ip2IsNil {
		return false
	}

	if len(ip1) != len(ip2) {
		return false
	}

	for i, val1 := range ip1 {
		val2 := ip2[i]

		if val1 != val2 {
			return false
		}
	}

	return true
}

func convertIpAddressToString(ip IpAddress) string {
	segmentStrs := make([]string, len(ip))
	for i, seg := range ip {
		segmentStrs[i] = fmt.Sprintf("%d", seg)
	}

	return strings.Join(segmentStrs, ".")
}
