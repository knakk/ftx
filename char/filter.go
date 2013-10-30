// Package char contains functions for manipulating strings on a character by
// by character basis.
package char

import "strings"

var removeGraphs = strings.NewReplacer("!", "", `\`, "", "#", "", "$", "", "%",
	"", "&", "", "'", "", "(", "", ")", "", "*", "", "+", "", ",", "", "-", "",
	".", "", "/", "", ":", "", ";", "", "<", "", "=", "", ">", "", "?", "", "@",
	"", "[", "", `\`, "", "]", "", "^", "", "`", "", "{", "", "|", "", "}", "",
	"~", "", ")", "", `"`, "")

var graphs2Space = strings.NewReplacer("!", " ", `\`, " ", "#", " ", "$", " ",
	"%", " ", "&", " ", "'", " ", "(", " ", ")", " ", "*", " ", "+", " ", ",",
	" ", "-", " ", ".", " ", "/", " ", ":", " ", ";", " ", "<", " ", "=", " ",
	">", " ", "?", " ", "@", " ", "[", " ", `\`, " ", "]", " ", "^", " ", "`",
	" ", "{", " ", "|", " ", "}", " ", "~", " ", ")", " ", `"`, " ")

var transformNorwegianChars = strings.NewReplacer(
	// Transform most of Latin1 subset to ASCII equivialents:
	"À", "A", "Á", "A", "Â", "A", "Ã", "A",
	"Ç", "C",
	"È", "E", "É", "E", "Ê", "E", "Ë", "E",
	"Ì", "I", "Í", "I", "Î", "I", "Ï", "I",
	"Ñ", "N",
	"Ò", "O", "Ó", "O", "Ô", "O", "Õ", "O",
	"Ù", "U", "Ú", "U", "Û", "U", "Ü", "U",
	"Ý", "Y",
	"à", "a", "á", "a", "â", "a", "ã", "a",
	"ç", "c", "ć", "c",
	"è", "e", "é", "e", "ê", "e", "ë", "e",
	"ì", "i", "í", "i", "î", "i", "ï", "i",
	"ñ", "n",
	"ò", "o", "ó", "o", "ô", "o", "õ", "o",
	"ù", "u", "ú", "u", "û", "u", "ü", "u",
	"ý", "y", "ÿ", "y",
	// Swedish -> Norwegian:
	"Ä", "Æ", "ä", "æ", "Ö", "Ø", "ö", "ø")

// Filter is an inteface for character filters ment to be applied before a
// string is passed to the tokenizer.
type Filter interface {
	Filter(s string) string
}

// ReplaceFilter replaces a list of strings with replacements.
type ReplaceFilter struct {
	replacePatterns *strings.Replacer
}

// NewPunct2SpaceFilter returns a ReplaceFilter which replaces punctuation
// and graphs characters with space.
func NewPunct2SpaceFilter() *ReplaceFilter {
	return &ReplaceFilter{graphs2Space}
}

// NewRemovePunctuationFilter returns a ReplaceFilter which removes punctuation
// and graph characters.
func NewRemovePunctuationFilter() *ReplaceFilter {
	return &ReplaceFilter{removeGraphs}
}

// NewNorwegianFoldingFilter returns a ReplaceFilter which works almost like
// an ASCII folding filter, but keeps æøå characters. It also converts
// Swedish characters ö -> ø and ä -> æ.
func NewNorwegianFoldingFilter() *ReplaceFilter {
	return &ReplaceFilter{transformNorwegianChars}
}

// Filter replaces a given regex pattern with a character.
func (f *ReplaceFilter) Filter(s string) string {
	return f.replacePatterns.Replace(s)
}

type StripHTMLFilter struct{} // TODO
