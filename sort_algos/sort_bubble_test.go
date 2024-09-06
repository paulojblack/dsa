package sort_algos

import (
	"fmt"
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, tc := range TestCases {
		sorted := bubbleSort(tc.in)
		if slices.Equal(sorted, tc.want) == false {
			t.Errorf("Bubble: %q, want: %q, got: %q", tc.in, tc.want, sorted)
		} else {
			fmt.Printf("Bubble: %q passed. Matches testcase: %q \n", tc.in, tc.want)
		}
	}

}
