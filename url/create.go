package url

import (
	"fmt"
	net_url "net/url"
)

func CreateUrlForIp(ip string, port uint, endpoint string, protocol string) string {
	u := net_url.URL{}

	u.Scheme = protocol
	u.Host = fmt.Sprintf("%s:%d", ip, port)
	u.Path = endpoint

	return u.String()
}
