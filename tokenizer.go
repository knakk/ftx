package ftx

import (
	"bytes"
	"text/scanner"
	"unicode/utf8"
)

const whiteSpace = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<'\f' | 1<<' '

// Tokenizer interface. A tokenizer takes a string and splits it into a slice
// of strings (tokens).
type Tokenizer interface {
	Tokenize(s string) []string
}

// WhiteSpaceTokenizer tokenize by splitting on whitespace.
type WhiteSpaceTokenizer struct{}

// NewWhiteSpaceTokenizer returns a new whitespace tokenizer.
func NewWhiteSpaceTokenizer() *WhiteSpaceTokenizer {
	return &WhiteSpaceTokenizer{}
}

// Tokenize splits a string at whitespaces.
func (t *WhiteSpaceTokenizer) Tokenize(s string) []string {
	words := []string{}
	var sc scanner.Scanner
	var token bytes.Buffer
	sc.Init(bytes.NewBufferString(s))
	sc.Mode = 0
	sc.Whitespace = 0
	charCount := 0
	r := sc.Scan()

	for r != scanner.EOF {
		if whiteSpace&(1<<uint(r)) != 0 {
			if charCount != 0 {
				words = append(words, token.String())
			}
			token.Reset()
			charCount = 0
		} else {
			token.WriteRune(r)
			charCount++
		}
		r = sc.Scan()
	}
	words = append(words, token.String())
	return words
}

// NGramTokenizer tokenize by n-grams
type NGramTokenizer struct {
	Min, Max int // codepoint size of a single n-gram
}

// NewNGramTokenizer returns a NGram tokenizer with min & max gramsize.
func NewNGramTokenizer(min, max int) *NGramTokenizer {
	if min < 1 { // min must be greater than zero
		min = 1
	}
	if min > max { // min must not be greater than max
		min, max = max, min
	}
	return &NGramTokenizer{min, max}
}

// Tokenize splits a strings into n-grams
func (t *NGramTokenizer) Tokenize(s string) []string {
	results := []string{}
	words := []string{}
	var sc scanner.Scanner
	var token bytes.Buffer
	sc.Init(bytes.NewBufferString(s))
	sc.Mode = 0
	sc.Whitespace = 0
	charCount := 0
	r := sc.Scan()

	for r != scanner.EOF {
		if whiteSpace&(1<<uint(r)) != 0 {
			if charCount != 0 {
				words = append(words, token.String())
			}
			token.Reset()
			charCount = 0
		} else {
			token.WriteRune(r)
			charCount++
		}
		r = sc.Scan()
	}
	words = append(words, token.String())
	for _, w := range words {
		blen := len(w)
		for c := range w {
			i := t.Min
			for c+i <= blen {
				if !utf8.ValidString(w[c : c+i]) {
					i++
					continue
				}
				rlen := len([]rune(w[c : c+i]))
				if rlen > t.Max {
					break
				}
				if rlen >= t.Min {
					results = append(results, w[c:c+i])
				}
				i++
			}
		}
	}
	return results
}
