package leetcode

import (
	"sort"
)

// 148. Sort List
// Sort a linked list in O(n log n) time using constant space complexity.

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	s := []int{}

	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}

	sort.Ints(s)

	var fresh = &ListNode{}
	var node = fresh
	for _, j := range s {
		node.Next = &ListNode{Val: j}
		node = node.Next
	}

	return fresh.Next
}
