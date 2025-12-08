package lookup

func Country(msisdn string) string {
	msisdn = normalize(msisdn)

	// probaj duže prefikse prvo
	maxLen := maxCountryPrefixLen
	if maxLen == 0 {
		maxLen = 3
	}
	for length := maxLen; length >= 1; length-- {
		if len(msisdn) >= length {
			prefix := msisdn[:length]
			if country, ok := countryByPrefix[prefix]; ok {
				return country
			}
		}
	}
	return "Unknown"
}
