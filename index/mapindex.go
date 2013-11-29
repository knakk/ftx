package index

import (
	"sync"

	"github.com/knakk/intset"
)

// MapIndex is the simplest index type. It cannot rank, but only records the
// occourences of tokens in documents.
type MapIndex struct {
	index map[string]*intset.HashSet
	sync.RWMutex
}

// NewMapIndex returns a new MapIndex.
func NewMapIndex() *MapIndex {
	return &MapIndex{index: make(map[string]*intset.HashSet)}
}

// CanRank states that MapIndex cannot rank.
func (i *MapIndex) CanRank() bool {
	return false
}

// Add indexes tokens by a give id in the index.
func (i *MapIndex) Add(doc int, tokens []string) {
	i.Lock()
	defer i.Unlock()
	for _, t := range tokens {
		if _, ok := i.index[t]; !ok {
			i.index[t] = intset.NewHashSet(999)
		}
		i.index[t].Add(doc)
	}
}

// Size returns the number of tokens in the index
func (i *MapIndex) Size() int {
	i.RLock()
	defer i.RUnlock()
	return len(i.index)
}

// Remove removes occurences of tokens in doc to the index.
func (i *MapIndex) Remove(doc int, tokens []string) {
	i.Lock()
	defer i.Unlock()
	for _, t := range tokens {
		i.index[t].Remove(doc)
		if i.index[t].Size() == 0 {
			delete(i.index, t)
		}
	}
}

// Query the MapIndex for search hits.
func (i *MapIndex) Query(q *Query) *SearchResults {
	res := SearchResults{}
	var and, not, or, setRes *intset.HashSet

	i.RLock()

	for _, t := range q.MustMatch {
		if _, ok := i.index[t]; ok {
			if and == nil {
				and = i.index[t].Clone()
			} else {
				and = intset.NewHashSet(999).Intersection(i.index[t])
			}
		} else {
			and = intset.NewHashSet(999)
			break
		}
	}

	for _, t := range q.MustNotMatch {
		if _, ok := i.index[t]; ok {
			if not == nil {
				not = i.index[t].Clone()
			} else {
				not = intset.NewHashSet(999).Intersection(i.index[t])
			}
		} else {
			not = intset.NewHashSet(999)
			break
		}
	}

	// Ignore q.ShouldMatch if q.MustMatch has any entries
	if and == nil {
		for _, t := range q.ShouldMatch {
			if _, ok := i.index[t]; ok {
				if or == nil {
					or = i.index[t].Clone()
				} else {
					or = intset.NewHashSet(999).Union(i.index[t])
				}
			}
		}
		setRes = or.Difference(not)
	} else {
		setRes = and.Difference(not)
	}

	i.RUnlock() // done reading from the index

	for i := range setRes.All() {
		res.Hits = append(res.Hits, searchHit{i, 0})
	}

	res.Count = setRes.Size()
	return &res
}
