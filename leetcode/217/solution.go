package leetcode

import "sort"

// 217. Contains Duplicate
// Given an array of integers, find if the array contains any duplicates.
// Your function should return true if any value appears at least twice in the array,
// and it should return false if every element is distinct.p

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i, j := range nums {
		if i == 0 {
			continue
		}

		if j == nums[i-1] {
			return true
		}
	}
	return false
}
