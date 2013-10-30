package ftx

import (
	"github.com/knakk/ftx/char"
	"github.com/knakk/ftx/index"
	"github.com/knakk/ftx/token"
)

// Analyzer defines how you index a document, by declaring which filters,
// tokenizer and index type you'll use.
type Analyzer struct {
	charFilters  []char.Filter
	tokenizer    Tokenizer
	tokenFilters []token.Filter
	Idx          index.Index
}

// Index indexes a given document using the character filters, tokenizer and
// token filters in the analyzer.
func (a Analyzer) Index(doc string, id int) {
	for _, f := range a.charFilters {
		doc = f.Filter(doc)
	}
	tokens := a.tokenizer.Tokenize(doc)
	for _, t := range a.tokenFilters {
		tokens = t.Filter(tokens)
	}
	a.Idx.Add(id, tokens)
}

func NewStandardAnalyzer() *Analyzer {
	return &Analyzer{
		charFilters: []char.Filter{
			char.NewRemovePunctuationFilter(),
			char.NewNorwegianFoldingFilter()},
		tokenizer:    NewWhiteSpaceTokenizer(),
		tokenFilters: []token.Filter{token.NewLowerCaseFilter()},
		Idx:          index.NewMapIndex(),
	}
}

func NewNGramAnalyzer(min, max int) *Analyzer {
	return &Analyzer{
		charFilters: []char.Filter{
			char.NewRemovePunctuationFilter(),
			char.NewNorwegianFoldingFilter()},
		tokenizer:    NewNGramTokenizer(min, max),
		tokenFilters: []token.Filter{token.NewLowerCaseFilter()},
		Idx:          index.NewMapIndex(),
	}
}
