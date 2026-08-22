package config

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func ParseConfigFromFlags() (Config, error) {
	fromIpStr := flag.String("from", "192.168.0.1", "The IP V4 address to start scanning from")
	toIpStr := flag.String("to", "192.168.0.100", "The IP V4 address to end scanning at (including this IP)")
	port := flag.Uint("port", 80, "The port to send requests to during the scan")
	endpoint := flag.String("endpoint", "", "The endpoint to send requests to during the scan")
	concurrentRequests := flag.Uint("concurrentrequests", 32, "How many concurrent HTTP requests should be sent out at once?")
	textToFind := flag.String("searchtext", "hello world!", "The text that we want to find in a HTTP response body")

	flag.Parse()

	fromIp, err := parseIpV4String(*fromIpStr)
	if err != nil {
		return Config{}, err
	}

	toIp, err := parseIpV4String(*toIpStr)
	if err != nil {
		return Config{}, err
	}

	config := Config{
		FromIp:             fromIp,
		ToIp:               toIp,
		Port:               *port,
		Endpoint:           *endpoint,
		ConcurrentRequests: *concurrentRequests,
		TextToFind:         *textToFind,
	}

	return config, nil
}

func parseIpV4String(ipStr string) ([]byte, error) {
	segmentStrs := strings.Split(ipStr, ".")
	if len(segmentStrs) != 4 {
		return []byte{}, fmt.Errorf("ip address '%s' should have 4 segments", ipStr)
	}

	segmentBytes := make([]byte, 0, 4)

	for _, segmentStr := range segmentStrs {
		segNum, err := strconv.ParseUint(segmentStr, 10, 8)
		if err != nil {
			return []byte{}, fmt.Errorf("ip address segment '%s' in ip address '%s' is not valid", segmentStr, ipStr)
		}

		segmentBytes = append(segmentBytes, byte(segNum))
	}

	return segmentBytes, nil
}
