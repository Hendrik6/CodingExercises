package leetcode

import "fmt"

// 1295. Find Numbers with Even Number of Digits
// Given an array nums of integers, return how many of them contain an even number of digits.

func findNumbers(nums []int) int {
	var evenDigits int
	for i := range nums {
		if len(fmt.Sprintf("%d", nums[i]))%2 == 0 {
			evenDigits++
		}
	}
	return evenDigits
}
