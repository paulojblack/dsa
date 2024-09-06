package sort_algos

var TestCases = []struct {
	in   []int
	want []int
}{
	{[]int{5, 2, 1, 7, 3, 10, 3, 5, 4}, []int{1, 2, 3, 3, 4, 5, 5, 7, 10}},
	{[]int{2, 5}, []int{2, 5}},
	{[]int{5, 2, 1, 7, 3, 10, 5, 4}, []int{1, 2, 3, 4, 5, 5, 7, 10}},
	{[]int{1, 2, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 2}},
}
