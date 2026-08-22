package main

import (
	"findip/config"
	"fmt"
	"os"
)

func main() {
	config, err := config.ParseConfigFromFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("From IP: %d.%d.%d.%d\n", config.FromIp[0], config.FromIp[1], config.FromIp[2], config.FromIp[3])
	fmt.Printf("To IP: %d.%d.%d.%d\n", config.ToIp[0], config.ToIp[1], config.ToIp[2], config.ToIp[3])
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Endpoint: '%s'\n", config.Endpoint)
	fmt.Printf("Concurrent requests: %d\n", config.ConcurrentRequests)
	fmt.Printf("Search Text: '%s'\n", config.TextToFind)
}
