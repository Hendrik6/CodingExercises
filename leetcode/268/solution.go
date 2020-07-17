package leetcode

// 268. Missing Number
//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
func missingNumber(nums []int) int {
	ans := len(nums)
	for i, num := range nums {
		ans ^= i ^ num
	}
	return ans
}
