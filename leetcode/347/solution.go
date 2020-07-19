package leetcode

import "sort"

//347. Top K Frequent Elements
//Given a non-empty array of integers, return the k most frequent elements.

func topKFrequent(nums []int, k int) []int {
	tmp, temp, res := make(map[int]int), [][]int{}, []int{}
	for _, v := range nums {
		tmp[v]++
	}
	for i, v := range tmp {
		temp = append(temp, []int{i, v})
	}
	sort.Slice(temp, func(a, b int) bool { return temp[a][1] > temp[b][1] })
	for i := 0; i < k; i++ {
		res = append(res, temp[i][0])
	}
	return res
}
