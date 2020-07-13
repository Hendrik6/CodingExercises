package leetcode

// 1. Two Sum
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}
		index[b] = i
	}
	return nil
}
