package ftx

import (
	"testing"
)

func BenchmarkWhiteSpaceTokenizer(b *testing.B) {
	t := NewWhiteSpaceTokenizer()
	for i := 0; i < b.N; i++ {
		t.Tokenize("Et støkke \"tekst\" å splitte opp i biter, versågod!")
	}
}
func BenchmarkNGramTokenizer(b *testing.B) {
	t := NewNGramTokenizer(2, 3)
	for i := 0; i < b.N; i++ {
		t.Tokenize("Et støkke \"tekst\" å splitte opp i biter, versågod!")
	}
}

func TestWhiteSpaceTokenizer(t *testing.T) {
	tokenizer := NewWhiteSpaceTokenizer()
	str := "go  ahead\nand, tokenize	me."
	tokens := tokenizer.Tokenize(str)
	want := []string{"go", "ahead", "and,", "tokenize", "me."}
	for i, tok := range tokens {
		if tok != want[i] {
			t.Fatalf("WhiteSpaceTokenizer.Tokenize(%q) => %#v; want %#v", str, tokens, want)
		}
	}
}

func TestNGramTokenizer(t *testing.T) {
	tokenizer := NewNGramTokenizer(2, 3)
	str := " FC \tSchølke\n04"
	tokens := tokenizer.Tokenize(str)
	want := []string{"FC", "Sch", "Sc", "chø", "ch", "høl", "hø", "ølk", "øl", "lke", "lk", "ke", "04"}
	if len(tokens) != len(want) {
		t.Fatalf("NGramTokenizer.Tokenize(%q) => %#v; want %#v", str, tokens, want)
	}
	for i, tok := range tokens {
		if tok != want[i] {
			t.Fatalf("NGramTokenizer.Tokenize(%q) => %#v; want %#v", str, tokens, want)
		}
	}
}
