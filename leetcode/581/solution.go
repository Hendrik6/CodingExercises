package leetcode

import "math"

// 581. Shortest Unsorted Continuous Subarray
// Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order,
// then the whole array will be sorted in ascending order, too.
// You need to find the shortest such subarray and output its length.

func findUnsortedSubarray(nums []int) int {
	start, end, min, max := 0, -1, math.MaxInt32, math.MinInt32
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
		if nums[j] <= min {
			min = nums[j]
		} else {
			start = j
		}
	}
	return end - start + 1
}
