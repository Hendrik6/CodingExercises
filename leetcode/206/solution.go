package leetcode

//206. Reverse Linked List
// Reverse a singly linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//Making a new node
	var front *ListNode

	//creating two variables which are both the same head param
	mid, end := head, head

	//if head isn't nil
	for mid != nil {

		//set the end as the next node
		end = mid.Next

		mid.Next = front
		front, mid = mid, end
	}
	return front
}
