package scan_test

import (
	"errors"
	"findip/config"
	"findip/ip"
	"findip/scan"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"testing"
	"time"
)

func TestRunIpScanWillCallCorrectIpsInBatches(t *testing.T) {
	expectedBatchSize := 5

	// setup mocks
	originalHttpGet := scan.HttpGet
	defer func() {
		scan.HttpGet = originalHttpGet
	}()

	// add data to this to allow the current batch's requests to finish (the amount of additions should match the length
	// of the current batch)
	getRespondTriggerChan := make(chan int, 15)

	callUrls := make([]string, 0, expectedBatchSize) // the URLs that HttpGet has been called with

	scan.HttpGet = func(url string) (*http.Response, error) {
		callUrls = append(callUrls, url)

		<-getRespondTriggerChan

		resp := &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
		}

		return resp, nil
	}

	// run scan
	testConfig := config.Config{
		FromIp:     ip.IpAddress{192, 168, 0, 1},
		ToIp:       ip.IpAddress{192, 168, 0, 15},
		Port:       80,
		Endpoint:   "",
		Protocol:   "http",
		BatchSize:  uint(expectedBatchSize),
		TextToFind: "Test text",
	}

	go func() {
		scan.RunIpScan(testConfig)
	}()

	// Check batch 1/3
	waitForBatchToMakeAllRequests()

	expectedUrls1 := []string{
		"http://192.168.0.1:80",
		"http://192.168.0.2:80",
		"http://192.168.0.3:80",
		"http://192.168.0.4:80",
		"http://192.168.0.5:80",
	}

	testCallUrlsAgainstExpected(t, callUrls, expectedUrls1, 1)

	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	callUrls = callUrls[:0] // length to 0

	// Check batch 2/3
	waitForBatchToMakeAllRequests()

	expectedUrls2 := []string{
		"http://192.168.0.6:80",
		"http://192.168.0.7:80",
		"http://192.168.0.8:80",
		"http://192.168.0.9:80",
		"http://192.168.0.10:80",
	}

	testCallUrlsAgainstExpected(t, callUrls, expectedUrls2, 2)

	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	callUrls = callUrls[:0]

	// Check batch 3/3
	waitForBatchToMakeAllRequests()

	expectedUrls3 := []string{
		"http://192.168.0.11:80",
		"http://192.168.0.12:80",
		"http://192.168.0.13:80",
		"http://192.168.0.14:80",
		"http://192.168.0.15:80",
	}

	testCallUrlsAgainstExpected(t, callUrls, expectedUrls3, 3)

	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
}

func TestRunIpScanWillConstructCorrectURLs(t *testing.T) {
	// setup mocks
	originalHttpGet := scan.HttpGet
	defer func() {
		scan.HttpGet = originalHttpGet
	}()

	callUrls := make([]string, 0, 5) // the URLs that HttpGet has been called with

	scan.HttpGet = func(url string) (*http.Response, error) {
		callUrls = append(callUrls, url)

		resp := &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
		}

		return resp, nil
	}

	// run scan
	tests := []struct {
		name         string
		config       config.Config
		expectedUrls []string
	}{
		{
			name: "standard",
			config: config.Config{
				FromIp:     ip.IpAddress{192, 168, 2, 1},
				ToIp:       ip.IpAddress{192, 168, 2, 5},
				Port:       80,
				Endpoint:   "",
				Protocol:   "http",
				BatchSize:  uint(5),
				TextToFind: "Test text",
			},
			expectedUrls: []string{
				"http://192.168.2.1:80",
				"http://192.168.2.2:80",
				"http://192.168.2.3:80",
				"http://192.168.2.4:80",
				"http://192.168.2.5:80",
			},
		},
		{
			name: "port",
			config: config.Config{
				FromIp:     ip.IpAddress{192, 168, 0, 1},
				ToIp:       ip.IpAddress{192, 168, 0, 5},
				Port:       81,
				Endpoint:   "",
				Protocol:   "http",
				BatchSize:  uint(5),
				TextToFind: "Test text",
			},
			expectedUrls: []string{
				"http://192.168.0.1:81",
				"http://192.168.0.2:81",
				"http://192.168.0.3:81",
				"http://192.168.0.4:81",
				"http://192.168.0.5:81",
			},
		},
		{
			name: "endpoint",
			config: config.Config{
				FromIp:     ip.IpAddress{192, 168, 2, 1},
				ToIp:       ip.IpAddress{192, 168, 2, 5},
				Port:       80,
				Endpoint:   "/test",
				Protocol:   "http",
				BatchSize:  uint(5),
				TextToFind: "Test text",
			},
			expectedUrls: []string{
				"http://192.168.2.1:80/test",
				"http://192.168.2.2:80/test",
				"http://192.168.2.3:80/test",
				"http://192.168.2.4:80/test",
				"http://192.168.2.5:80/test",
			},
		},
		{
			name: "protocol",
			config: config.Config{
				FromIp:     ip.IpAddress{192, 168, 2, 1},
				ToIp:       ip.IpAddress{192, 168, 2, 5},
				Port:       80,
				Endpoint:   "",
				Protocol:   "https",
				BatchSize:  uint(5),
				TextToFind: "Test text",
			},
			expectedUrls: []string{
				"https://192.168.2.1:80",
				"https://192.168.2.2:80",
				"https://192.168.2.3:80",
				"https://192.168.2.4:80",
				"https://192.168.2.5:80",
			},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scan.RunIpScan(tt.config)

			testCallUrlsAgainstExpected(t, callUrls, tt.expectedUrls, i+1)

			callUrls = callUrls[:0] // length to 0
		})
	}
}

func TestRunIpScanWillRunCorrectAmountOfBatches(t *testing.T) {
	// setup mocks
	originalHttpGet := scan.HttpGet
	defer func() {
		scan.HttpGet = originalHttpGet
	}()

	// add data to this to allow the current batch's requests to finish (the amount of additions should match the length
	// of the current batch)
	getRespondTriggerChan := make(chan int, 9)

	callCount := 0

	scan.HttpGet = func(url string) (*http.Response, error) {
		callCount++

		<-getRespondTriggerChan

		resp := &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
		}

		return resp, nil
	}

	// run
	testConf := config.Config{
		FromIp:     ip.IpAddress{192, 168, 2, 1},
		ToIp:       ip.IpAddress{192, 168, 2, 9},
		Port:       80,
		Endpoint:   "",
		Protocol:   "http",
		BatchSize:  uint(3),
		TextToFind: "Test text",
	}

	go func() {
		scan.RunIpScan(testConf)
	}()

	// we expected 3 batches of 3
	waitForBatchToMakeAllRequests()
	if callCount != 3 {
		t.Errorf("expected total calls after batch 1 to be 3. got %d", callCount)
	}

	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	waitForBatchToMakeAllRequests()
	if callCount != 6 {
		t.Errorf("expected total calls after batch 2 to be 6. got %d", callCount)
	}

	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	waitForBatchToMakeAllRequests()
	if callCount != 9 {
		t.Errorf("expected total calls after batch 3 to be 9. got %d", callCount)
	}

	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
	getRespondTriggerChan <- 1
}

func TestRunIpScanWillIgnoreErrorsFromHttpRequest(t *testing.T) {
	// setup mocks
	originalHttpGet := scan.HttpGet
	defer func() {
		scan.HttpGet = originalHttpGet
	}()

	callCount := 0

	scan.HttpGet = func(url string) (*http.Response, error) {
		callCount++

		// Failure during the second batch
		if callCount == 5 {
			return nil, errors.New("test error to be ignored")
		}

		resp := &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
		}

		return resp, nil
	}

	// run
	testConf := config.Config{
		FromIp:     ip.IpAddress{192, 168, 2, 1},
		ToIp:       ip.IpAddress{192, 168, 2, 9},
		Port:       80,
		Endpoint:   "",
		Protocol:   "http",
		BatchSize:  uint(3),
		TextToFind: "Test text",
	}

	scan.RunIpScan(testConf)

	if callCount != 9 {
		t.Errorf("expected total calls count to be 9. got %d", callCount)
	}
}

func TestRunIpScanWillIgnoreBadResponseFromHttpRequest(t *testing.T) {
	// setup mocks
	originalHttpGet := scan.HttpGet
	defer func() {
		scan.HttpGet = originalHttpGet
	}()

	callCount := 0

	scan.HttpGet = func(url string) (*http.Response, error) {
		callCount++

		// Failure during the second batch
		if callCount == 5 {
			resp := &http.Response{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
			}

			return resp, nil
		}

		resp := &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
		}

		return resp, nil
	}

	// run
	testConf := config.Config{
		FromIp:     ip.IpAddress{192, 168, 2, 1},
		ToIp:       ip.IpAddress{192, 168, 2, 9},
		Port:       80,
		Endpoint:   "",
		Protocol:   "http",
		BatchSize:  uint(3),
		TextToFind: "Test text",
	}

	scan.RunIpScan(testConf)

	if callCount != 9 {
		t.Errorf("expected batch 3 calls count to be 9. got %d", callCount)
	}
}

func TestRunIpScanWillReturnFoundIp(t *testing.T) {
	// setup mocks
	originalHttpGet := scan.HttpGet
	defer func() {
		scan.HttpGet = originalHttpGet
	}()

	callCount := 0
	searchText := "this is the text you're looking for"
	successUrl := "http://192.168.2.5:80"

	scan.HttpGet = func(url string) (*http.Response, error) {
		callCount++

		// Find the correct text, during the second batch
		if url == successUrl {
			bodyText := fmt.Sprintf("guess what? %s. Yep, that's right!", searchText)

			resp := &http.Response{
				StatusCode: 200,
				Status:     "200 OK",
				Body:       io.NopCloser(strings.NewReader(bodyText)),
			}

			return resp, nil
		}

		resp := &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
		}

		return resp, nil
	}

	// run
	testConf := config.Config{
		FromIp:     ip.IpAddress{192, 168, 2, 1},
		ToIp:       ip.IpAddress{192, 168, 2, 9},
		Port:       80,
		Endpoint:   "",
		Protocol:   "http",
		BatchSize:  uint(3),
		TextToFind: searchText,
	}

	result := scan.RunIpScan(testConf)

	// assert
	expectedIp := "192.168.2.5"
	if result != expectedIp {
		t.Errorf("expected returned ip to be '%s'. got '%s'", expectedIp, result)
	}

	if callCount != 6 {
		t.Errorf("expected scan to stop with the second batch. got %d http calls", callCount)
	}
}

func TestRunIpScanWillReturnEmptyStringIfSearchTextNotFound(t *testing.T) {
	// setup mocks
	originalHttpGet := scan.HttpGet
	defer func() {
		scan.HttpGet = originalHttpGet
	}()

	scan.HttpGet = func(url string) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader("this is not the ip you're looking for, move along")),
		}

		return resp, nil
	}

	// run
	testConf := config.Config{
		FromIp:     ip.IpAddress{192, 168, 2, 1},
		ToIp:       ip.IpAddress{192, 168, 2, 9},
		Port:       80,
		Endpoint:   "",
		Protocol:   "http",
		BatchSize:  uint(3),
		TextToFind: "Test text",
	}

	result := scan.RunIpScan(testConf)

	if result != "" {
		t.Errorf("expected result to be an empty string. got '%s'", result)
	}
}

func testCallUrlsAgainstExpected(t *testing.T, callUrls []string, expectedUrls []string, batchNum int) {
	if len(callUrls) != len(expectedUrls) {
		t.Errorf("expected batch %d to have %d calls. got %d", batchNum, len(expectedUrls), len(callUrls))
		return
	}

	sortedCallUrls := make([]string, len(callUrls))
	copy(sortedCallUrls, callUrls)
	slices.Sort(sortedCallUrls)
	noDuplicateSortedCallUrls := slices.Compact(sortedCallUrls)

	if len(noDuplicateSortedCallUrls) != len(expectedUrls) {
		t.Errorf("expected batch %d to have %d unique urls. got %d", batchNum, len(expectedUrls), len(noDuplicateSortedCallUrls))
		return
	}

	for _, url := range noDuplicateSortedCallUrls {
		if !slices.Contains(expectedUrls, url) {
			t.Errorf("url '%s' in batch %d was not expected", url, batchNum)
		}
	}
}

func waitForBatchToMakeAllRequests() {
	// Simplistic solution, but fine for our purposes
	time.Sleep(2 * time.Second)
}
