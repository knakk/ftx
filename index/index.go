// Package index contains the index interface and some implementations.
package index

// Index interface all indexes must implement.
type Index interface {
	Add(doc int, tokens []string)
	Remove(doc int, tokens []string)
	CanRank() bool
	Query(q *Query) *SearchResults
	Size() int // how many tokens are indexed
}

// Query represent a boolean query. Allowed combinations: AND+NOT, OR+NOT
type Query struct {
	MustMatch    []string // Boolean AND
	MustNotMatch []string // Boolean NOT
	ShouldMatch  []string // Boolean OR (has no effect if MustMatch is used)
	Limit        int      // 0 = no limit
}

// NewQuery returns a new boolean query.
func NewQuery() *Query {
	return &Query{
		make([]string, 0),
		make([]string, 0),
		make([]string, 0),
		0}
}

// Must taktes terms that must occour to constitute a search hit.
func (q *Query) Must(tokens []string) *Query {
	for _, t := range tokens {
		q.MustMatch = append(q.MustMatch, t)
	}
	return q
}

// Should takes terms that should occour in search hit, i.e you'll at least
// term occourse to get a hit.
func (q *Query) Should(tokens []string) *Query {
	for _, t := range tokens {
		q.ShouldMatch = append(q.ShouldMatch, t)
	}
	return q
}

// Not takes terms that should not occour in search hit.
func (q *Query) Not(tokens []string) *Query {
	for _, t := range tokens {
		q.MustNotMatch = append(q.MustNotMatch, t)
	}
	return q
}

type searchHit struct {
	ID    int
	Score float64 // allways set to 0 if the index isn't ranked
}

// SearchResults stores the search hits.
type SearchResults struct {
	Count int
	Hits  []searchHit
}

//func (sr *SearchResults) SortById() *SearchResults {}
//func (sr *SearchResults) SortByRank() *SearchResults{}
//func (sr *SearchResults) Reverse() *SearchResults
