package token

import (
	"strings"
	"testing"

	"github.com/knakk/specs"
)

func TestLowerCase(t *testing.T) {
	s := specs.New(t)

	tokens := strings.Split("Gi MÆ di TÅKENAN", " ")

	f := NewLowerCaseFilter()
	tokens = f.Filter(tokens)

	tests := []specs.Spec{
		{"gi", tokens[0]},
		{"mæ", tokens[1]},
		{"di", tokens[2]},
		{"tåkenan", tokens[3]},
	}
	s.ExpectAll(tests)
}
