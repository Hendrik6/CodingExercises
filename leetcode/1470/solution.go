package leetcode

// 1470. Shuffle the Array
// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
// Return the array in the form [x1,y1,x2,y2,...,xn,yn].

func shuffle(nums []int, n int) []int {
	var ret []int

	for i, j := range nums {
		if len(ret) == len(nums) {
			break
		}
		ret = append(ret, j)
		ret = append(ret, nums[i+n])
	}

	return ret
}
