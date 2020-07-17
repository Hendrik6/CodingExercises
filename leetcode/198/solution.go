package leetcode

// 198. House Robber
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed
// the only constraint stopping you from robbing each of them is
// that adjacent houses have security system connected and
// it will automatically contact the police if two adjacent houses were broken into on the same night.
// Given a list of non-negative integers representing the amount of money of each house,
// determine the maximum amount of money you can rob tonight without alerting the police.

func rob(nums []int) int {
	p1 := 0
	p2 := 0
	for _, v := range nums {
		tmp := p1
		p1 = max(v+p2, p1)
		p2 = tmp
	}
	return p1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
