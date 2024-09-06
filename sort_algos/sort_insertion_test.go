package sort_algos

import (
	"fmt"
	"slices"
	"testing"
)

func TestInsertionSortAsc(t *testing.T) {
	for _, tc := range TestCases {
		sorted := insertionSortAsc(tc.in)
		if slices.Equal(sorted, tc.want) == false {
			t.Errorf("insertionOrderAc: %q, want: %q", tc.in, tc.want)
		} else {
			fmt.Printf("insertionOrderAsc: %v passed. Matches testcase: %v \n", tc.in, tc.want)
		}
	}

}

func TestInsertionSortDesc(t *testing.T) {

	for _, tc := range TestCases {
		sorted := insertionSortDesc(tc.in)
		if slices.Equal(sorted, tc.want) == false {
			t.Errorf("insertionOrderDesc: %q, want: %q", tc.in, tc.want)
		} else {
			fmt.Printf("insertionOrderDesc: %v passed. Matches testcase: %v \n", tc.in, tc.want)
		}
	}

}
