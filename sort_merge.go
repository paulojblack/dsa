package main

func mergeSort(nums []int) []int {
	return mergeRecurse(nums)
}

func mergeRecurse(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	midpoint := len(nums) / 2

	left := mergeRecurse(nums[:midpoint])
	right := mergeRecurse(nums[midpoint:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
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
