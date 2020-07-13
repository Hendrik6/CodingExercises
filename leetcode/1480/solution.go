package leetcode

//1480. Running Sum of 1d Array
//Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//Return the running sum of nums.

func runningSum(nums []int) []int {
	var sum []int
	for _, j := range nums {
		if len(sum) != 0 {
			j = j + sum[len(sum)-1]
		}
		sum = append(sum, j)
	}
	return sum
}
