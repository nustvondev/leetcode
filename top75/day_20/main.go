package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Remove Nth Node From End of List
// Technique: Two pointers (fast & slow)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast := dummy
	slow := dummy

	// Move fast n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both until fast reaches end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove target node
	slow.Next = slow.Next.Next
	return dummy.Next
}

func buildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for _, v := range arr[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := buildList([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 2)
	printList(head) // 1 2 3 5

	head2 := buildList([]int{1})
	head2 = removeNthFromEnd(head2, 1)
	printList(head2) // []

	head3 := buildList([]int{1, 2})
	head3 = removeNthFromEnd(head3, 1)
	printList(head3) // 1
}
