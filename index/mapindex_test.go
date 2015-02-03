package index

import (
	"testing"

	"github.com/knakk/intset"
)

func TestMapIndexAddDoc(t *testing.T) {
	tests := []struct {
		doc      int
		tokens   []string
		wantSize int
	}{
		{1, []string{"alle"}, 1},
		{1, []string{"kan", "ikke", "synge"}, 4},
		{1, []string{"kan"}, 4},
		{2, []string{"kan"}, 4},
	}
	idx := NewMapIndex()
	for _, tt := range tests {
		idx.Add(tt.doc, tt.tokens)
		if idx.Size() != tt.wantSize {
			t.Errorf("mapIndex.Size() => %v; want %v", idx.Size(), tt.wantSize)
		}
	}
}

func TestMapIndexRemoveDoc(t *testing.T) {
	idx := NewMapIndex()
	idx.Add(1, []string{"alle", "kan", "ikke", "synge"})
	if idx.Size() != 4 {
		t.Errorf("idx.Size() => %d; want 4", idx.Size())
	}
	if idx.Size() != 4 {
		t.Errorf("idx.Size() => %d; want 4", idx.Size())
	}
	idx.Remove(2, []string{"ikke"})
	if idx.Size() != 4 {
		t.Errorf("idx.Size() => %d; want 4", idx.Size())
	}
	idx.Remove(1, []string{"ikke"})
	if idx.Size() != 3 {
		t.Errorf("idx.Size() => %d; want 3", idx.Size())
	}

}

func srAsIntSet(sr *SearchResults) *intset.SliceSet {
	s := intset.NewSliceSet(100)
	for _, h := range sr.Hits {
		s.Add(h.ID)
	}
	return s
}

func TestMapIndexQuery(t *testing.T) {
	idx := NewMapIndex()
	idx.Add(1, []string{"alle", "nonner", "drar", "skjønt"})
	idx.Add(12, []string{"skjønt", "alle", "må", "noe"})
	idx.Add(88, []string{"alle", "kan", "ikke"})
	idx.Add(9, []string{"belgskveppe", "æeaåedår"})

	q := NewQuery().Must([]string{"alle", "skjønt"}).Not([]string{"kan"})
	res := idx.Query(q)
	if res.Count != 2 {
		t.Errorf("expected 2 results, got %d", res.Count)
	}
	if !srAsIntSet(res).Contains(1, 12) {
		t.Errorf("got %v; expected (1,12)", res)
	}

	q2 := NewQuery().Should([]string{"æeaåedår", "alle"}).Not([]string{"nonner"})
	res2 := idx.Query(q2)
	if !srAsIntSet(res2).Contains(12, 88, 9) {
		t.Errorf("got %v; expected (12,88,9)", res2)
	}

	q3 := NewQuery().Must([]string{"belgskveppe", "all"})
	res3 := idx.Query(q3)
	if res3.Count != 0 {
		t.Errorf("got %v, want no results", res3)
	}
	//s.Expect(res3.Hits[0], "hva?")
}
