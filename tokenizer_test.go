package ftx

import (
	"testing"

	"github.com/knakk/specs"
)

func TestWhiteSpaceTokenizer(t *testing.T) {
	s := specs.New(t)

	tokenizer := NewWhiteSpaceTokenizer()
	str := "go  ahead\nand, tokenize	me."
	tokens := tokenizer.Tokenize(str)
	tests := []specs.Spec{
		{tokens[0], "go"},
		{tokens[1], "ahead"},
		{tokens[2], "and,"},
		{tokens[3], "tokenize"},
		{tokens[4], "me."},
	}
	s.ExpectAll(tests)
}

func TestNGramTokenizer(t *testing.T) {
	s := specs.New(t)

	tokenizer := NewNGramTokenizer(2, 3)
	str := "FC Schølke 04"
	tokens := tokenizer.Tokenize(str)
	tests := []specs.Spec{
		{tokens[0], "FC"},
		{tokens[1], "Sc"},
		{tokens[2], "Sch"},
		{tokens[3], "ch"},
		{tokens[4], "chø"},
		{tokens[5], "hø"},
		{tokens[6], "høl"},
		{tokens[7], "øl"},
		{tokens[8], "ølk"},
		{tokens[9], "lk"},
		{tokens[10], "lke"},
		{tokens[11], "ke"},
		{tokens[12], "04"},
	}
	s.ExpectAll(tests)
}
