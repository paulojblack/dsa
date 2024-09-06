package sort_algos

import (
	"fmt"
	"slices"
	"testing"
)

func TestMergeSort(t *testing.T) {

	for _, tc := range TestCases {
		sorted := mergeSort(tc.in)
		if slices.Equal(sorted, tc.want) == false {
			t.Errorf("Merge: %q, want: %q, got: %q", tc.in, tc.want, sorted)
		} else {
			fmt.Printf("Merge: %q passed. Matches testcase: %q \n", tc.in, tc.want)
		}
	}

}
