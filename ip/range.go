package ip

import (
	"fmt"
	"strings"
)

func GenerateIpAddressStrings(fromIp IpAddress, toIp IpAddress) []string {
	if IpAddressesAreEqual(fromIp, toIp) {
		return []string{convertIpAddressToString(fromIp)}
	}

	currentIp := make(IpAddress, len(fromIp))
	copy(currentIp, fromIp)

	ipsToScan := make([]string, 0, 1024) // capacity here is just a high value, to save on memory reshuffling in most cases

	lastIpHasBeenProcessed := false
	for !lastIpHasBeenProcessed {
		incrementIpAddress(&currentIp)

		ipStr := convertIpAddressToString(currentIp)
		ipsToScan = append(ipsToScan, ipStr)

		lastIpHasBeenProcessed = IpAddressesAreEqual(currentIp, toIp)
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

func convertIpAddressToString(ip IpAddress) string {
	segmentStrs := make([]string, len(ip))
	for i, seg := range ip {
		segmentStrs[i] = fmt.Sprintf("%d", seg)
	}

	return strings.Join(segmentStrs, ".")
}
