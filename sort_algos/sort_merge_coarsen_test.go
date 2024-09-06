package sort_algos

import (
	"fmt"
	"slices"
	"testing"
)

func TestMergeSortCoarsen(t *testing.T) {
	k := 5

	for _, tc := range TestCases {
		sorted := mergeSortCoarsen(tc.in, k)
		if slices.Equal(sorted, tc.want) == false {
			t.Errorf("Merge Coarsen: %q, want: %q, got: %q", tc.in, tc.want, sorted)
		} else {
			fmt.Printf("Merge Coarsen: %q passed. Matches testcase: %q \n", tc.in, tc.want)
		}
	}

}
