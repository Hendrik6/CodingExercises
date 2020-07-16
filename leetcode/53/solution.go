package leetcode

// 53. Maximum Subarray
// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.

func maxSubArray(nums []int) int {
	var sum int
	var max int
	max = nums[0]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
