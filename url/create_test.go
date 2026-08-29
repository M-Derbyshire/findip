package url_test

import (
	"findip/url"
	"testing"
)

func TestCreateUrlForIpWillBuildTheCorrectIP(t *testing.T) {
	tests := []struct {
		name        string
		ip          string
		port        uint
		endpoint    string
		protocol    string
		expectedUrl string
	}{
		{
			name:        "http",
			ip:          "192.168.0.1",
			port:        80,
			endpoint:    "/test1",
			protocol:    "http",
			expectedUrl: "http://192.168.0.1:80/test1",
		},
		{
			name:        "https",
			ip:          "192.168.0.2",
			port:        80,
			endpoint:    "/test2",
			protocol:    "https",
			expectedUrl: "https://192.168.0.2:80/test2",
		},
		{
			name:        "port",
			ip:          "192.168.0.3",
			port:        101,
			endpoint:    "/test3",
			protocol:    "http",
			expectedUrl: "http://192.168.0.3:101/test3",
		},
		{
			name:        "empty-endpoint",
			ip:          "192.168.0.1",
			port:        80,
			endpoint:    "",
			protocol:    "http",
			expectedUrl: "http://192.168.0.1:80",
		},
		{
			name:        "single-slash-endpoint",
			ip:          "192.168.0.1",
			port:        80,
			endpoint:    "/",
			protocol:    "http",
			expectedUrl: "http://192.168.0.1:80/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := url.CreateUrlForIp(tt.ip, tt.port, tt.endpoint, tt.protocol)

			if result != tt.expectedUrl {
				t.Errorf("expected url to be '%s'. got '%s'", tt.expectedUrl, result)
			}
		})
	}
}
