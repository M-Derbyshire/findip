package config

type Config struct {
	FromIp             []byte `description:"The IP V4 address to start scanning from. Each byte represents a number in the address (e.g. 192 in 192.168.0.1)"`
	ToIp               []byte `description:"The IP V4 address to end scanning at (including this IP). Each byte represents a number in the address (e.g. 192 in 192.168.0.1)"`
	Port               uint   `description:"The port to send requests to during the scan"`
	Endpoint           string `description:"The endpoint to send requests to during the scan"`
	ConcurrentRequests uint   `description:"How many concurrent HTTP requests should be sent out at once?"`
	TextToFind         string `description:"The text that we want to find in a HTTP response body"`
}
