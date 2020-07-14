package leetcode

// 1512. Number of Good Pairs

// Given an array of integers nums.
// A pair (i,j) is called good if nums[i] == nums[j] and i < j.
// Return the number of good pairs.

func numIdenticalPairs(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				ret++
			}
		}
	}
	return ret
}
