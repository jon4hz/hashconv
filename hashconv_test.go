package hashconv

import (
	"testing"
)

type testList []struct {
	name, got, exp string
}

func TestHashes(t *testing.T) {
	testList{
		{"hashes(0)", Format(0), "0 H"},
		{"hashes(1)", Format(1), "1 H"},
		{"hashes(803)", Format(803), "803 H"},
		{"hashes(999)", Format(999), "999 H"},

		{"hashes(1024)", Format(1024), "1.024 kH"},
		{"hashes(9999)", Format(9999), "9.999 kH"},
		{"hashes(1MH - 1)", Format(MHash - Hash), "999.999 kH"},

		{"hashes(1024 * 1024)", Format(1024 * 1024), "1.049 MH"},
		{"hashes(1GH - 1K)", Format(GHash - KHash), "999.999 MH"},

		{"hashes(1GH)", Format(GHash), "1 GH"},
		{"hashes(1TH - 1M)", Format(THash - MHash), "999.999 GH"},
		{"hashes(9999 * 1000)", Format(9999 * 1000), "9.999 MH"},

		{"hashes(1TH)", Format(THash), "1 TH"},
		{"hashes(1PH - 1T)", Format(PHash - THash), "999 TH"},

		{"hashes(1PH)", Format(PHash), "1 PH"},
		{"hashes(1PH - 1T)", Format(EHash - PHash), "999 PH"},

		{"hashes(1EH)", Format(EHash), "1 EH"},

		{"hashes(5.5GH)", Format(5.5 * GHash), "5.5 GH"},
	}.validate(t)
}

func BenchmarkHashs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Format(16.5 * GHash)
	}
}

func (tl testList) validate(t *testing.T) {
	for _, test := range tl {
		if test.got != test.exp {
			t.Errorf("On %v, expected '%v', but got '%v'",
				test.name, test.exp, test.got)
		}
	}
}
