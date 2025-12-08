package lookup

import "strings"

func NumberType(msisdn string) string {
	// msisdn = strings.TrimPrefix(msisdn, "+")
	msisdn = normalize(msisdn)

	// prepoznaj Italiju
	if strings.HasPrefix(msisdn, "39") {
		rest := msisdn[2:] // posle "39"
		if strings.HasPrefix(rest, "3") {
			return "mobile"
		}
		return "fixed"
	}

	// prepoznaj Srbiju
	if strings.HasPrefix(msisdn, "381") {
		rest := msisdn[3:] // posle "381"
		if strings.HasPrefix(rest, "6") {
			return "mobile"
		}
		return "fixed"
	}

	// Switzerland
	if strings.HasPrefix(msisdn, "41") {
		rest := msisdn[2:]
		if strings.HasPrefix(rest, "7") {
			return "mobile"
		}
		return "fixed"
	}

	// Greece
	if strings.HasPrefix(msisdn, "30") {
		rest := msisdn[2:]
		if strings.HasPrefix(rest, "69") {
			return "mobile"
		}
		return "fixed"
	}

	return "unknown"
}
