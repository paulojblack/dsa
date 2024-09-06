package main

import (
	"math"
)

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
	totalLength := len(left) + len(right)
	result := make([]int, totalLength)
	leftCopy := append([]int{}, left...)
	leftCopy = append(leftCopy, math.MaxInt)

	rightCopy := append([]int{}, right...)
	rightCopy = append(rightCopy, math.MaxInt)
	l := 0
	r := 0
	for k := 0; k < totalLength; k++ {
		if leftCopy[l] <= rightCopy[r] {
			result[k] = leftCopy[l]
			l++
		} else {
			result[k] = rightCopy[r]
			r++
		}
	}
	return result
}
