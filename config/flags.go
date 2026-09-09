package config

import (
	"errors"
	"findip/config/behaviourflag"
	"findip/ip"
	"flag"
)

func ParseConfigFromFlags() (Config, error) {
	// Config for scan
	fromIpStr := flag.String("from", "192.168.0.1", "The IP V4 address to start scanning from")
	toIpStr := flag.String("to", "192.168.0.100", "The IP V4 address to end scanning at (including this IP)")
	port := flag.Uint("port", 80, "The port to send requests to during the scan")
	endpoint := flag.String("endpoint", "", "The endpoint to send requests to during the scan")
	protocol := flag.String("protocol", "http", "The protocol to use in requests (e.g. http or https)")
	batchSize := flag.Uint("batchsize", 32, "How many concurrent HTTP requests should be sent out at once?")
	textToFind := flag.String("searchtext", "hello world!", "The text that we want to find in a HTTP response body")

	// Check requested behaviour
	var printVersionFlag behaviourflag.PrintVersionFlag
	flag.Var(&printVersionFlag, "version", printVersionFlag.Usage())
	flag.Var(&printVersionFlag, "v", printVersionFlag.Usage())

	// Parse and validate
	flag.Parse()

	fromIp, err := ip.ParseIpV4String(*fromIpStr)
	if err != nil {
		return Config{}, err
	}

	toIp, err := ip.ParseIpV4String(*toIpStr)
	if err != nil {
		return Config{}, err
	}

	if !ip.IpAddressesAreEqual(fromIp, toIp) {
		ipsRightWayRound, err := ip.IpAddressIsLessThan(fromIp, toIp)
		if err != nil {
			return Config{}, err
		}

		if !ipsRightWayRound {
			return Config{}, errors.New("-from ip address is greater than -to ip address")
		}
	}

	requestedBehaviour := BehaviourScan
	if printVersionFlag.IsSet() {
		requestedBehaviour = BehaviourPrintVersion
	}

	// Construct config
	config := Config{
		FromIp:             fromIp,
		ToIp:               toIp,
		Port:               *port,
		Endpoint:           *endpoint,
		Protocol:           *protocol,
		BatchSize:          *batchSize,
		TextToFind:         *textToFind,
		RequestedBehaviour: requestedBehaviour,
	}

	return config, nil
}
