package sort_algos

var threshold int

func mergeSortCoarsen(nums []int, k int) []int {
	threshold = k
	return mergeRecurseCoarsen(nums)
}

func mergeRecurseCoarsen(nums []int) []int {
	if len(nums) <= threshold {
		return lilInsertion(nums)
	}
	midpoint := len(nums) / 2

	left := mergeRecurseCoarsen(nums[:midpoint])
	right := mergeRecurseCoarsen(nums[midpoint:])

	return mergeCoarsen(left, right)
}

func mergeCoarsen(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	totalLength := len(left) + len(right)
	// Walkers
	l, r := 0, 0

	for l+r < totalLength {
		if r >= len(right) {
			for l < len(left) {
				result[l+r] = left[l]
				l++
			}
			break
		}
		if l >= len(left) {
			for r < len(right) {
				result[l+r] = right[r]
				r++
			}
			break
		}

		if left[l] > right[r] {
			result[l+r] = right[r]
			r++
		} else {
			result[l+r] = left[l]
			l++
		}
	}

	return result
}

func lilInsertion(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			nums[j] = key
			j--
		}
	}
	return nums
}
