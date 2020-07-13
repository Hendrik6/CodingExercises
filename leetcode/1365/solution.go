package leetcode

// 1365. How Many Numbers Are Smaller Than the Current Number
// Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
// That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].
// Return the answer in an array.

func smallerNumbersThanCurrent(nums []int) []int {
	var arr []int
	for _, j := range nums {
		counter := 0
		for _, i := range nums {
			if i < j {
				counter++
			}
		}
		arr = append(arr, counter)

	}
	return arr
}
