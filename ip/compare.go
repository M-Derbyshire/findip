package ip

func IpAddressesAreEqual(ip1 IpAddress, ip2 IpAddress) bool {
	ip1IsNil := ip1 == nil
	ip2IsNil := ip2 == nil

	if ip1IsNil != ip2IsNil {
		return false
	}

	if len(ip1) != len(ip2) {
		return false
	}

	for i, val1 := range ip1 {
		val2 := ip2[i]

		if val1 != val2 {
			return false
		}
	}

	return true
}
