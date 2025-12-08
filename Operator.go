package lookup

var operatorByPrefix = map[string]string{
	"38160": "MTS (original range)",
	"38161": "A1 (original range)",
	"38163": "Yettel (original range)",
	"38162": "Yettel (original range)",
	"38165": "MTS (original range)",
	"38164": "MTS (original range)",
	"38166": "MTS (original range)",
	"3069":  "Greek mobile (Cosmote/Vodafone/WIND)",
	"30231": "Greek fixed (OTE - Thessaloniki)",
	"30210": "Greek fixed (OTE - Athens)",
}

var maxOperatorPrefixLen int

func init() {
	for prefix := range operatorByPrefix {
		if l := len(prefix); l > maxOperatorPrefixLen {
			maxOperatorPrefixLen = l
		}
	}
}

func Operator(msisdn string) string {
	msisdn = normalize(msisdn)

	maxLen := maxOperatorPrefixLen
	if maxLen == 0 {
		maxLen = 5
	}

	for length := maxLen; length >= 1; length-- {
		if len(msisdn) >= length {
			prefix := msisdn[:length]
			if op, ok := operatorByPrefix[prefix]; ok {
				return op
			}
		}
	}

	return "Unknown"
}
