package ftx

import (
	"testing"

	"github.com/knakk/ftx/index"
	"github.com/knakk/intset"
	"github.com/knakk/specs"
)

func srAsIntSet(sr *index.SearchResults) intset.IntSet {
	s := intset.New()
	for _, h := range sr.Hits {
		s.Add(h.ID)
	}
	return s
}

func TestStandardAnalyzer(t *testing.T) {
	s := specs.New(t)
	a := NewStandardAnalyzer()

	a.Index("Kan du finne Meg?", 1)
	a.Index("æøå DU di dæ DÅ", 2)
	a.Index("han hun du meg deg", 3)

	q := index.NewQuery().Must([]string{"meg"})
	res := a.index.Query(q)
	s.Expect(srAsIntSet(res).Contains(1, 3), true)
}

func TestNGramAnalyzer(t *testing.T) {
	s := specs.New(t)
	a := NewNGramAnalyzer(2, 10)

	a.Index("Bokstavlig talt!", 2)
	a.Index("KAKE-BOKSEN-SMULDRER-SMULER", 4)
	a.Index("Krepsens vendekrets", 8)
	a.Index("krapyl", 10)

	q := index.NewQuery().Should([]string{"oks", "kr"})
	res := a.index.Query(q)
	s.Expect(srAsIntSet(res).Contains(2, 4, 8, 10), true)

	q2 := index.NewQuery().Must([]string{"bok"}).Not([]string{"smuler"})
	res2 := a.index.Query(q2)
	s.Expect(srAsIntSet(res2).Equal(intset.NewFromSlice([]int{2})), true)
}
