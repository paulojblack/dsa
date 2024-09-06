package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{5, 2, 1, 7, 3, 10, 3, 5, 4}, []int{1, 2, 3, 3, 4, 5, 5, 7, 10}},
		{[]int{2, 5}, []int{2, 5}},
		{[]int{5, 2, 1, 7, 3, 10, 5, 4}, []int{1, 2, 3, 4, 5, 5, 7, 10}},
		{[]int{1, 2, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 2}},
	}

	for _, tc := range testCases {
		sorted := quickSort(tc.in)
		if slices.Equal(sorted, tc.want) == false {
			t.Errorf("Merge: %q, want: %q, got: %q", tc.in, tc.want, sorted)
		} else {
			fmt.Printf("Merge: %q passed. Matches testcase: %q \n", tc.in, tc.want)
		}
	}

}
