// Package token contains functions for transforming token streams.
package token

import "strings"

// Filter is an interface for filters that can modify, remove or add tokens
// in a token stream.
type Filter interface {
	Filter(tokens []string) []string
}

// LowerCaseFilter makes all tokens lower case.
type LowerCaseFilter struct{}

// NewLowerCaseFilter returns a new LowerCase token filter.
func NewLowerCaseFilter() *LowerCaseFilter {
	return &LowerCaseFilter{}
}

// Filter returns a stream of lower cased tokens.
func (t *LowerCaseFilter) Filter(tokens []string) []string {
	results := []string{}
	for _, t := range tokens {
		results = append(results, strings.ToLower(t))
	}
	return results
}
