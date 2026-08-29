package scan

import (
	"bytes"
	"findip/config"
	"findip/ip"
	"findip/url"
	"findip/util"
	"io"
	"net/http"
	"sync"
)

// Scans for the IP address that responds with a body that includes the expected search-text string. Returns
// an empty string if the IP is not found, otherwise returns the IP address as a string
var RunIpScan = func(conf config.Config) string {
	fullIpList := ip.GenerateIpAddressStrings(conf.FromIp, conf.ToIp)
	ipLists := util.SplitSliceIntoMaxLengthSlices(fullIpList, conf.BatchSize)

	for _, list := range ipLists {
		responses := sendRequests(list, conf.Port, conf.Endpoint, conf.Protocol)

		for _, resp := range responses {
			if responseContainsSearchText(resp.resp, conf.TextToFind) {
				return resp.ip
			}
		}
	}

	return ""
}

// Triggers the http requests to the given IPs. Ignores any errors/failures, and
// returns the responses from any requests that didn't return an error
func sendRequests(ips []string, port uint, endpoint string, protocol string) []responseData {
	maxCapacity := len(ips)
	responseChannel := make(chan responseData, maxCapacity)
	responses := make([]responseData, 0, maxCapacity)

	var wg sync.WaitGroup
	for _, ip := range ips {
		wg.Add(1)

		urlStr := url.CreateUrlForIp(ip, port, endpoint, protocol)

		go func() {
			defer wg.Done()

			resp, err := http.Get(urlStr)
			if err == nil {
				data := responseData{
					ip:   ip,
					resp: resp,
				}

				responseChannel <- data
			}
		}()
	}
	go func() {
		wg.Wait()
		close(responseChannel)
	}()

	for resp := range responseChannel {
		responses = append(responses, resp)
	}

	return responses
}

func responseContainsSearchText(resp *http.Response, searchText string) bool {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	return bytes.Contains(body, []byte(searchText))
}

type responseData struct {
	ip   string
	resp *http.Response
}
