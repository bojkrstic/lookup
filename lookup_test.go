package lookup

import "testing"

func TestCountry(t *testing.T) {
	cases := []struct {
		msisdn string
		want   string
	}{
		{"+393383260866", "Italy"},
		{"+38164123456", "Serbia"},
		{"+306941234567", "Greece"},
		{"+41071234567", "Switzerland"},
	}

	for _, tc := range cases {
		if got := Country(tc.msisdn); got != tc.want {
			t.Fatalf("Country(%s) = %s, want %s", tc.msisdn, got, tc.want)
		}
	}
}

func TestNumberType(t *testing.T) {
	cases := []struct {
		msisdn string
		want   string
	}{
		{"+393383260866", "mobile"},
		{"+390636918899", "fixed"},
		{"+38164123456", "mobile"},
		{"+38111345678", "fixed"},
		{"+306941234567", "mobile"},
		{"+302112345678", "fixed"},
	}

	for _, tc := range cases {
		if got := NumberType(tc.msisdn); got != tc.want {
			t.Fatalf("NumberType(%s) = %s, want %s", tc.msisdn, got, tc.want)
		}
	}
}

func TestIsValidLength(t *testing.T) {
	cases := []struct {
		msisdn string
		want   bool
	}{
		{"+393383260866", true},
		{"+38164123456", true},
		{"+390636", false},
		{"+3816", false},
		{"+306941234567", true},
		{"+3069", false},
	}

	for _, tc := range cases {
		if got := IsValidLength(tc.msisdn); got != tc.want {
			t.Fatalf("IsValidLength(%s) = %v, want %v", tc.msisdn, got, tc.want)
		}
	}
}

func TestOperator(t *testing.T) {
	cases := []struct {
		msisdn string
		want   string
	}{
		{"+381601234567", "MTS (original range)"},
		{"+381621234567", "A1 (original range)"},
		{"+381641234567", "Yettel (original range)"},
		{"+306912345678", "Greek mobile (Cosmote/Vodafone/WIND)"},
		{"+302310669985", "Greek fixed (OTE - Thessaloniki)"},
		{"+302109876543", "Greek fixed (OTE - Athens)"},
	}

	for _, tc := range cases {
		if got := Operator(tc.msisdn); got != tc.want {
			t.Fatalf("Operator(%s) = %s, want %s", tc.msisdn, got, tc.want)
		}
	}
}
