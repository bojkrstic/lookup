package lookup

var countryByPrefix = map[string]string{
	"1":   "USA/Canada",
	"30":  "Greece",
	"33":  "France",
	"39":  "Italy",
	"381": "Serbia",
	"385": "Croatia",
	"41":  "Switzerland",
}

var maxCountryPrefixLen int

func init() {
	for prefix := range countryByPrefix {
		if l := len(prefix); l > maxCountryPrefixLen {
			maxCountryPrefixLen = l
		}
	}
}

type LookupResponse struct {
	MSISDN      string `json:"msisdn"`
	Country     string `json:"country"`
	NumberType  string `json:"numberType"`
	ValidLength bool   `json:"validLength"`
	Operator    string `json:"operator"`
}
