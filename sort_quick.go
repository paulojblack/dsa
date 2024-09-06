package main

func quickSort(nums []int) []int {
	quickilydickily(nums, 0, len(nums)-1)
	return nums
}

func quickilydickily(nums []int, l, r int) {
	if l < r {
		p := partition(nums, l, r)
		quickilydickily(nums, l, p-1)
		quickilydickily(nums, p+1, r)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[r]

	slow := l - 1

	for l < r {
		if nums[l] <= pivot {
			slow++

			tmp := nums[l]
			nums[l] = nums[slow]
			nums[slow] = tmp
		}
		l++
	}
	slow++
	nums[r] = nums[slow]
	nums[slow] = pivot

	return slow
}
