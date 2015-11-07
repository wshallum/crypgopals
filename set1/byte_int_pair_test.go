package set1

import "testing"

func TestByteIntPairSorting(t *testing.T) {
	a := []byteIntPair{
		{b: 0, i: 10},
		{b: 10, i: 0},
	}

	sortByteIntPairs(a, byteIntPairSortByByteDesc)
	if a[0].b != 10 {
		t.Error("Fail")
	}
	sortByteIntPairs(a, byteIntPairSortByByteAsc)
	if a[0].b != 0 {
		t.Error("Fail")
	}
	sortByteIntPairs(a, byteIntPairSortByIntDesc)
	if a[0].i != 10 {
		t.Error("Fail")
	}
	sortByteIntPairs(a, byteIntPairSortByIntAsc)
	if a[0].i != 0 {
		t.Error("Fail")
	}
}
