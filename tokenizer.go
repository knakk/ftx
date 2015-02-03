package ftx

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

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
	bwords := bytes.Fields([]byte(s))
	for _, w := range bwords {
		words = append(words, string(w))
	}
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
func (t *NGramTokenizer) Tokenize(input string) []string {
	results := []string{}
	start := 0                // start of current word
	pos := 0                  // position in input string
	c := 0                    // count of characters in current word
	w := make([]int, t.Max+1) // keep track of rune widths
	var r rune
	var atEnd bool
	for {
		if !atEnd {
			r, w[c] = utf8.DecodeRuneInString(input[pos:])
			if w[c] == 0 {
				// reached end of string
				atEnd = true
			} else {
				pos += w[c]
				c++
			}
		}
		if unicode.IsSpace(r) {
			// reached word boundary
			if c >= t.Min {
				j := 0
				for i := c - 1; i >= t.Min; i-- {
					j += w[i]
					results = append(results, input[start:pos-j])
				}
			}
			// reset
			start = pos
			c = 0
			for i := range w {
				w[i] = 0
			}
			continue
		}
		if c == t.Max || atEnd {
			// reached max gram length or at the end of string
			j := 0
			for i := c; i >= t.Min; i-- {
				j += w[i]
				results = append(results, input[start:pos-j])
			}
			if start+t.Min >= len(input) {
				break
			}

			// advance one rune
			start += w[0]
			c--
			copy(w[0:], w[1:])
		}

	}
	return results
}
