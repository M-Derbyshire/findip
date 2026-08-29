package main

import (
	"findip/config"
	"findip/scan"
	"fmt"
	"os"
)

func main() {
	conf, err := config.ParseConfigFromFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	resultIp := scan.RunIpScan(conf)

	if resultIp == "" {
		fmt.Fprintln(os.Stderr, "unable to find ip")
		return
	}

	fmt.Println(resultIp)
}
