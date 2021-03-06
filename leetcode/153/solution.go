package leetcode

// 153. Find Minimum in Rotated Sorted Array
// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
// Find the minimum element.
// You may assume no duplicate exists in the array.

func findMin(nums []int) int {
	val := nums[0]

	for _, value := range nums {
		if value < val {
			val = value
		}
	}
	return val
}
