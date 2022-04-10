package unionfind

import (
	"reflect"
	"testing"
)

func TestUnionFind(t *testing.T) {
	set := make([]uint32, 7)
	for i := range set {
		Add(set, uint32(i))
	}

	// Connected components: {0,1,2}, {3,4,5}, {6}
	Union(set, 0, 1)
	Union(set, 1, 2)
	Union(set, 5, 4)
	Union(set, 4, 3)
	Union(set, 6, 6)

	// Find should return the min vertex from the connected component
	expect := []uint32{0, 0, 0, 3, 3, 3, 6}

	actual := make([]uint32, len(set))
	for i := range set {
		actual[i] = Find(set, uint32(i))
	}

	if !reflect.DeepEqual(actual, expect) {
		t.Error("unexpected result")
		t.Logf("actual: %v", actual)
		t.Logf("expect: %v", expect)
	}
}

func BenchmarkUnionFind(b *testing.B) {
	const count = 10_000

	b.Run("increasing-pairs", func(b *testing.B) {
		set := make([]uint32, count)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			zero(set)

			for j := uint32(0); j < count-1; j++ {
				Add(set, j+1)
				Union(set, j, j+1)
			}

			for j := range set {
				set[j] = Find(set, uint32(j))
			}
		}
	})

	b.Run("decreasing-pairs", func(b *testing.B) {
		set := make([]uint32, count)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			zero(set)

			Add(set, count-1)
			for j := uint32(count - 1); j > 0; j-- {
				Add(set, j-1)
				Union(set, j-1, j)
			}

			for j := len(set) - 1; j >= 0; j-- {
				set[j] = Find(set, uint32(j))
			}
		}
	})
}

func zero(set []uint32) {
	for i := range set {
		set[i] = 0
	}
}
