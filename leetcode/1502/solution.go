package leetcode

import "sort"

// 1502. Can Make Arithmetic Progression From Sequence

// Given an array of numbers arr.
// A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.
// Return true if the array can be rearranged to form an arithmetic progression, otherwise, return false.

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 1 {
		return true
	}

	sort.Ints(arr)

	diff := arr[1] - arr[0]

	for i, j := range arr {
		if i == 0 {
			continue
		}
		if j-arr[i-1] != diff {
			return false
		}
	}

	return true
}
