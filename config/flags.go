package config

import (
	"errors"
	"findip/ip"
	"flag"
)

func ParseConfigFromFlags() (Config, error) {
	fromIpStr := flag.String("from", "192.168.0.1", "The IP V4 address to start scanning from")
	toIpStr := flag.String("to", "192.168.0.100", "The IP V4 address to end scanning at (including this IP)")
	port := flag.Uint("port", 80, "The port to send requests to during the scan")
	endpoint := flag.String("endpoint", "", "The endpoint to send requests to during the scan")
	concurrentRequests := flag.Uint("concurrentrequests", 32, "How many concurrent HTTP requests should be sent out at once?")
	textToFind := flag.String("searchtext", "hello world!", "The text that we want to find in a HTTP response body")

	flag.Parse()

	fromIp, err := ip.ParseIpV4String(*fromIpStr)
	if err != nil {
		return Config{}, err
	}

	toIp, err := ip.ParseIpV4String(*toIpStr)
	if err != nil {
		return Config{}, err
	}

	ipsRightWayRound, err := ip.IpAddressIsLessThan(fromIp, toIp)
	if err != nil {
		return Config{}, err
	}

	if !ipsRightWayRound {
		return Config{}, errors.New("-from ip address is greater than -to ip address")
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
