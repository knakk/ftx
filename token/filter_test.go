package token

import (
	"strings"
	"testing"
)

func TestLowerCase(t *testing.T) {
	input := strings.Split("Gi MÆ di TÅKENAN", " ")
	f := NewLowerCaseFilter()
	output := f.Filter(input)
	want := []string{"gi", "mæ", "di", "tåkenan"}
	for i, tok := range want {
		if tok != output[i] {
			t.Fatalf("LowerCaseFilter(%v) => %v; want %v", input, output, want)
		}
	}
}
