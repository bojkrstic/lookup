package lookup

import "strings"

func IsValidLength(msisdn string) bool {
	//msisdn = strings.TrimPrefix(msisdn, "+")
	msisdn = normalize(msisdn)
	length := len(msisdn)

	// Italy (39)
	if strings.HasPrefix(msisdn, "39") {
		// Tipično 11 cifara: 39 + 9
		return length == 11
	}

	// Serbia (381)
	if strings.HasPrefix(msisdn, "381") {
		// Tipično 11 ili 12 cifara
		return length == 11 || length == 12
	}

	// Switzerland (41)
	if strings.HasPrefix(msisdn, "41") {
		// 41 + 9 cifara = 11
		return length == 11
	}

	// Greece (30)
	if strings.HasPrefix(msisdn, "30") {
		// 30 + 10 cifara = 12
		return length == 12
	}

	return false
}
