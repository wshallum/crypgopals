package set1

import (
	"sort"
)

type byteIntPair struct {
	b byte
	i int
}

const (
	byteIntPairSortByByteAsc = iota
	byteIntPairSortByByteDesc
	byteIntPairSortByIntAsc
	byteIntPairSortByIntDesc
)

type byteIntPairSorter struct {
	pairs []byteIntPair
	by    int
}

func (s *byteIntPairSorter) Len() int {
	return len(s.pairs)
}

func (s *byteIntPairSorter) Swap(i, j int) {
	s.pairs[i], s.pairs[j] = s.pairs[j], s.pairs[i]
}

func (s *byteIntPairSorter) Less(i, j int) bool {
	if s.by == byteIntPairSortByByteAsc {
		return s.pairs[i].b < s.pairs[j].b
	} else if s.by == byteIntPairSortByByteDesc {
		return s.pairs[i].b > s.pairs[j].b
	} else if s.by == byteIntPairSortByIntAsc {
		return s.pairs[i].i < s.pairs[j].i
	} else if s.by == byteIntPairSortByIntDesc {
		return s.pairs[i].i > s.pairs[j].i
	}
	// should not happen
	panic("oops")
}

func sortByteIntPairs(pairs []byteIntPair, by int) {
	if by < byteIntPairSortByByteAsc || by > byteIntPairSortByIntDesc {
		return
	}
	sorter := &byteIntPairSorter{pairs: pairs, by: by}
	sort.Sort(sorter)
}
