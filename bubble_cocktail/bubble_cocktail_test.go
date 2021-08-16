package gsort

import (
	"testing"
	"sort-go/sort_test"
)

func TestSort(t *testing.T) {
	for _, test := range sort_test.Cases {
		res := Sort(test.Input)
		if !sort_test.IsEqualSlice(test.Output, res) {
			t.Errorf("sort(%v) = %v, want %v", test.Input, res, test.Output)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(sort_test.Cases[0].Input)
	}
}
