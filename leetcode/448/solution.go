package leetcode

import "math"

// 448. Find All Numbers Disappeared in an Array
// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
// Find all the elements of [1, n] inclusive that do not appear in this array.
// Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		index := int(math.Abs(float64(num))) - 1
		nums[index] = -1 * int(math.Abs(float64(nums[index])))
	}

	var numbers []int
	for i, num := range nums {
		if num > 0 {
			numbers = append(numbers, i+1)
		}
	}
	return numbers
}
