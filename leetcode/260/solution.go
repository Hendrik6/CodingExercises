package leetcode

// 260. Single Number III
// Given an array of numbers nums, in which exactly two elements appear only once
// and all the other elements appear exactly twice.
// Find the two elements that appear only once.

func singleNumber(nums []int) []int {
	m := make(map[int]int)
	var ret []int

	for _, value := range nums {
		m[value]++
	}

	for _, v := range nums {
        if m[v] == 1 {
			ret = append(ret, v)
		}
	}

	return ret
}