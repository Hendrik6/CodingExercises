package leetcode

// 154. Find Minimum in Rotated Sorted Array II
// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
// Find the minimum element.
// The array may contain duplicates.

func findMin(nums []int) int {
	firstVal := nums[0]

	for _, value := range nums {
		if value < firstVal {
			firstVal = value
		}
	}
	return firstVal
}
