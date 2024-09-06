package main

func quickSort(nums []int) []int {
	quickilydickily(nums, 0, len(nums)-1)
	return nums
}

func quickilydickily(nums []int, l, r int) {
	if l < r {
		pivot := partition(nums, l, r)
		quickilydickily(nums, l, pivot-1)
		quickilydickily(nums, pivot+1, r)
	}

}

func partition(nums []int, l, r int) int {
	p := nums[r]

	i := l - 1

	for l < r {
		if nums[l] <= p {
			i++
			tmp := nums[i]
			nums[i] = nums[l]
			nums[l] = tmp
		}
		l++
	}

	nums[r] = nums[i+1]
	nums[i+1] = p

	return i + 1
}
