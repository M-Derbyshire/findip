package main

import (
	"findip/config"
	"findip/info"
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

	if conf.RequestedBehaviour == config.BehaviourPrintVersion {
		fmt.Println(info.Version)
		return
	}

	resultIp := scan.RunIpScan(conf)

	if resultIp == "" {
		fmt.Fprintln(os.Stderr, "unable to find ip")
		return
	}

	fmt.Println(resultIp)
}
