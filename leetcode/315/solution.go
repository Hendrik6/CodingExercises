package leetcode

// 315. Count of Smaller Numbers After Self
// You are given an integer array nums and you have to return a new counts array.
// The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

func countSmaller(nums []int) []int {
	smaller := make([]int, len(nums))
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < v {
				smaller[i]++
			}
		}
	}
	return smaller
}
