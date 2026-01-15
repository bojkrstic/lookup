package lookup

var operatorByPrefix = map[string]string{
	// Serbia mobile ranges (RATEL allocations, trunk 06 removed)
	"38160": "A1 Serbia (original range)",
	"38161": "A1 Serbia (original range)",
	"38162": "Yettel Serbia (original range)",
	"38163": "Yettel Serbia (original range)",
	"38164": "Telekom Srbija (mts original range)",
	"38165": "Telekom Srbija (mts original range)",
	"38166": "Telekom Srbija (mts original range)",
	"38167": "Globaltel Serbia (MVNO range)",
	"38169": "A1 Serbia (additional range)",
	// Serbia fixed examples
	"38111": "Serbia fixed (Belgrade)",
	"38118": "Serbia fixed (Niš)",
	"38121": "Serbia fixed (Novi Sad)",

	// Italy mobile ranges
	"39320": "Wind Tre Italy (320 prefix)",
	"39328": "Wind Tre Italy (328 prefix)",
	"39329": "Wind Tre Italy (329 prefix)",
	"39330": "TIM Italy (330 prefix)",
	"39333": "TIM Italy (333 prefix)",
	"39335": "TIM Italy (335 prefix)",
	"39338": "TIM Italy (338 prefix)",
	"39339": "TIM Italy (339 prefix)",
	"39340": "Vodafone Italy (340 prefix)",
	"39345": "Vodafone Italy (345 prefix)",
	"39348": "Vodafone Italy (348 prefix)",
	"39349": "Vodafone Italy (349 prefix)",
	"39351": "Iliad Italy (351 prefix)",
	"39370": "PosteMobile / Fastweb (370 prefix)",
	"393":   "Italian mobile (3xx range)",

	// Italy fixed samples (leading 0 dropped in E.164)
	"3902":  "Italy fixed (Milan 02)",
	"3906":  "Italy fixed (Rome 06)",
	"39081": "Italy fixed (Naples 081)",
	"390":   "Italy fixed (other geographic ranges)",

	// Switzerland mobile ranges
	"4174": "Lycamobile Switzerland (074 prefix)",
	"4176": "Sunrise UPC Switzerland (076 prefix)",
	"4178": "Salt Switzerland (078 prefix)",
	"4179": "Swisscom Mobile (079 prefix)",
	"417":  "Switzerland mobile (07x range)",

	// Switzerland fixed ranges
	"4121": "Switzerland fixed (Lausanne/Vaud 21)",
	"4122": "Switzerland fixed (Geneva 22)",
	"4131": "Switzerland fixed (Bern 31)",
	"4141": "Switzerland fixed (Central Switzerland 41)",
	"4144": "Switzerland fixed (Zürich 44)",
	"4171": "Switzerland fixed (St. Gallen 71)",
	"412":  "Switzerland fixed (02x/04x range)",

	// Greece mobile ranges
	"30690": "WIND Hellas (690 prefix)",
	"30693": "WIND Hellas (693 prefix)",
	"30694": "Vodafone Greece (694 prefix)",
	"30695": "Vodafone Greece (695 prefix)",
	"30697": "Cosmote Greece (697 prefix)",
	"30698": "Cosmote Greece (698 prefix)",
	"3069":  "Greek mobile (Cosmote/Vodafone/WIND)",

	// Greece fixed ranges
	"30231": "Greek fixed (OTE - Thessaloniki)",
	"30221": "Greek fixed (OTE - Thessaly / Central)",
	"30210": "Greek fixed (OTE - Athens)",
	"302":   "Greek fixed (other cities)",
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
