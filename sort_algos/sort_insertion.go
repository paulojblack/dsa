package sort_algos

func insertionSortAsc(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]

		for j := i - 1; j >= 0; j-- {
			if nums[j] > key {
				nums[j+1] = nums[j]
				nums[j] = key
			}
		}
	}

	return nums
}

func insertionSortDesc(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]

		j := i - 1
		for j >= 0 && nums[j] < key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	return nums
}
