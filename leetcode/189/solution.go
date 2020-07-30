package leetcode

// 189. Rotate Array
// Given an array, rotate the array to the right by k steps, where k is non-negative.

func rotate(nums []int, k int) {
	n := len(nums)
	rotate := make([]int, n)
	for i := 0; i < n; i++ {
		rotate[(i+k)%n] = nums[i]
	}
	copy(nums, rotate)
}
