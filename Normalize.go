package lookup

import "strings"

func normalize(msisdn string) string {
	msisdn = strings.TrimSpace(msisdn)           // skloni razmake sa početka/kraja
	msisdn = strings.ReplaceAll(msisdn, " ", "") // skloni sve razmake u sredini

	// podrži + i 00 kao internacionalni prefiks
	msisdn = strings.TrimPrefix(msisdn, "+")
	if strings.HasPrefix(msisdn, "00") {
		msisdn = msisdn[2:]
	}

	return msisdn
}
