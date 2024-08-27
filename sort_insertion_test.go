package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestInsertionSortAsc(t *testing.T) {
	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{5, 2, 1, 7, 3, 10, 3, 5, 4}, []int{1, 2, 3, 3, 4, 5, 5, 7, 10}},
		{[]int{1, 2, 1, 1, 1}, []int{1, 1, 1, 1, 2}},
	}

	for _, tc := range testCases {
		sorted := insertionSortAsc(tc.in)
		if slices.Equal(sorted, tc.want) == false {
			t.Errorf("insertionOrderAc: %q, want: %q", tc.in, tc.want)
		} else {
			fmt.Printf("insertionOrderAsc: %v passed. Matches testcase: %v \n", tc.in, tc.want)
		}
	}

}
