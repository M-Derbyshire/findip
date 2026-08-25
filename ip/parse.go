package ip

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseIpV4String(ipStr string) (IpAddress, error) {
	segmentStrs := strings.Split(ipStr, ".")
	if len(segmentStrs) != 4 {
		return IpAddress{}, fmt.Errorf("ip address '%s' should have 4 segments", ipStr)
	}

	segmentBytes := make(IpAddress, 0, 4)

	for _, segmentStr := range segmentStrs {
		segNum, err := strconv.ParseUint(segmentStr, 10, 8)
		if err != nil {
			return IpAddress{}, fmt.Errorf("ip address segment '%s' in ip address '%s' is not valid", segmentStr, ipStr)
		}

		segmentBytes = append(segmentBytes, byte(segNum))
	}

	return segmentBytes, nil
}
