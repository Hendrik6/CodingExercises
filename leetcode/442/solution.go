package leetcode

// 442. Find All Duplicates in an Array
// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
// Find all the elements that appear twice in this array.
// Could you do it without extra space and in O(n) runtime?

func findDuplicates(nums []int) []int {
	m := make(map[int]bool)
	var rs []int

	for _, v := range nums {
		_, ok := m[v]
		if ok {
			rs = append(rs, v)
		} else {
			m[v] = true
		}
	}

	return rs
}
