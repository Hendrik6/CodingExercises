package leetcode

//136. Single Number
//Given a non-empty array of integers, every element appears twice except for one. Find that single one.

func singleNumber(nums []int) int {
	ret := 0
	for i := range nums {
		ret ^= nums[i]
	}
	return ret
}
