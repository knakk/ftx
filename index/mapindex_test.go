package index

import (
	"testing"

	"github.com/knakk/intset"
	"github.com/knakk/specs"
)

func TestMapIndexAddDoc(t *testing.T) {
	s := specs.New(t)

	idx := NewMapIndex()
	s.Expect(idx.Size(), 0)
	idx.Add(1, []string{"alle"})
	s.Expect(idx.Size(), 1)
	idx.Add(1, []string{"kan", "ikke", "synge"})
	s.Expect(idx.Size(), 4)
	idx.Add(1, []string{"kan"})
	s.Expect(idx.Size(), 4)
	idx.Add(2, []string{"kan"})
	s.Expect(idx.Size(), 4)
}

func TestMapIndexRemoveDoc(t *testing.T) {
	s := specs.New(t)

	idx := NewMapIndex()
	idx.Add(1, []string{"alle", "kan", "ikke", "synge"})
	s.Expect(idx.Size(), 4)
	idx.Add(2, []string{"kan"})
	s.Expect(idx.Size(), 4)
	idx.Remove(2, []string{"ikke"})
	s.Expect(idx.Size(), 4)
	idx.Remove(1, []string{"ikke"})
	s.Expect(idx.Size(), 3)
}

func srAsIntSet(sr *SearchResults) intset.IntSet {
	s := intset.New()
	for _, h := range sr.Hits {
		s.Add(h.ID)
	}
	return s
}

func TestMapIndexQuery(t *testing.T) {
	s := specs.New(t)

	idx := NewMapIndex()
	idx.Add(1, []string{"alle", "nonner", "drar", "skjønt"})
	idx.Add(12, []string{"skjønt", "alle", "må", "noe"})
	idx.Add(88, []string{"alle", "kan", "ikke"})
	idx.Add(9, []string{"belgskveppe", "æeaåedår"})

	q := NewQuery().Must([]string{"alle", "skjønt"}).Not([]string{"kan"})
	res := idx.Query(q)
	s.Expect(res.Count, 2)
	s.Expect(srAsIntSet(res).Contains(1, 12), true)

	q2 := NewQuery().Should([]string{"æeaåedår", "alle"}).Not([]string{"nonner"})
	res2 := idx.Query(q2)
	s.Expect(srAsIntSet(res2).Contains(12, 88, 9), true)
}
