package quicksort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tables := []struct {
		unsorted []int
		sorted   []int
	}{
		{
			[]int{1, 8, 3, 47, 2, 5, 7},
			[]int{1, 2, 3, 5, 7, 8, 47},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{4, 3, 2, 1},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{2, 1, 1, 1},
			[]int{1, 1, 1, 2},
		},
		{
			[]int{1, 2, 1, 1},
			[]int{1, 1, 1, 2},
		},
		{
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
		},
	}

	for _, table := range tables {
		sorted := make([]int, len(table.unsorted))
		copy(sorted, table.unsorted)
		Sort(sorted)

		if !reflect.DeepEqual(sorted, table.sorted) {
			t.Errorf("Sort of (%v) was incorrect, got: %v, want: %v.", table.unsorted, sorted, table.sorted)
		}
	}
}
