package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Iterative solution – O(n)
Idea:
- Use 3 pointers: prev, curr, next
- Reverse links step by step
*/
func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/*
Recursive solution – O(n)
Idea:
- Base case: head == nil or head.Next == nil → return head
- Recurse on head.Next, then reverse pointer
*/
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// helper to build linked list
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, v := range nums[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

// helper to convert linked list to slice
func toSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	list := buildList([]int{1, 2, 3, 4, 5})
	fmt.Println(toSlice(reverseListIterative(list))) // [5 4 3 2 1]

	list2 := buildList([]int{1, 2})
	fmt.Println(toSlice(reverseListRecursive(list2))) // [2 1]

	list3 := buildList([]int{})
	fmt.Println(toSlice(reverseListIterative(list3))) // []
}
