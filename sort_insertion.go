package main

func insertionSortAsc(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]

		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
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
