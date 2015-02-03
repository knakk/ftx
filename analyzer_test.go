package ftx

import (
	"testing"

	"github.com/knakk/ftx/index"
	"github.com/knakk/intset"
)

func srAsIntSet(sr *index.SearchResults) *intset.HashSet {
	s := intset.NewHashSet(100)
	for _, h := range sr.Hits {
		s.Add(h.ID)
	}
	return s
}

func TestStandardAnalyzer(t *testing.T) {
	a := NewStandardAnalyzer()

	a.Index("Kan du finne Meg?", 1)
	a.Index("æøå DU di dæ DÅ", 2)
	a.Index("han hun du meg deg", 3)

	q := index.NewQuery().Must([]string{"meg"})
	res := a.Idx.Query(q)
	if !srAsIntSet(res).Contains(1, 3) {
		t.Fatal("documents not indexed/not queryable")
	}

	q3 := index.NewQuery().Must([]string{"du", "DÆ", "han"})
	res3 := a.Idx.Query(q3)
	if res3.Count != 0 {
		t.Errorf("expected no results, got %v", res3)
	}

	a.UnIndex("Kan du finne Meg?", 1)
	res2 := a.Idx.Query(q)
	if srAsIntSet(res2).Contains(1) {
		t.Error("UnIndex didn't remove document")
	}
}

func TestNGramAnalyzer(t *testing.T) {
	a := NewNGramAnalyzer(2, 10)

	a.Index("Bokstavlig talt!", 2)
	a.Index("KAKE-BOKSEN-SMULDRER-SMULER", 4)
	a.Index("Krepsens vendekrets", 8)
	a.Index("krapyl", 10)

	q := index.NewQuery().Should([]string{"oks", "kr"})
	res := a.Idx.Query(q)
	if !srAsIntSet(res).Contains(2, 4, 8, 10) {
		t.Fatal("documents not indexed/not queryable")
	}

	q2 := index.NewQuery().Must([]string{"bok"}).Not([]string{"smuler"})
	res2 := a.Idx.Query(q2)
	if !srAsIntSet(res2).Equal(intset.NewHashSet(10).Add(2)) {
		t.Error("ngramanalyzer: must+not query fails")
	}
}
