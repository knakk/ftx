package ftx

import "testing"

func BenchmarkWhiteSpaceTokenizer(b *testing.B) {
	t := NewWhiteSpaceTokenizer()
	for i := 0; i < b.N; i++ {
		t.Tokenize("Et støkke tekst å splitte opp i biter, versågod!")
	}
}
