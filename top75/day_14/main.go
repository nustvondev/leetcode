package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Approach 1 – HashSet
- Traverse list, store visited nodes in a map.
- If a node is visited again → cycle detected.
- O(n) time, O(n) space
*/
func hasCycleHash(head *ListNode) bool {
	visited := map[*ListNode]bool{}
	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}

/*
Approach 2 – Floyd’s Cycle Detection (Tortoise & Hare)
- Use two pointers: slow and fast.
- Move slow by 1 step, fast by 2 steps.
- If cycle exists, they will eventually meet.
- If fast reaches nil → no cycle.
- O(n) time, O(1) space
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// helper to build linked list with optional cycle
func buildList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	var cycleNode *ListNode
	if pos == 0 {
		cycleNode = head
	}
	for i, v := range nums[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
		if i+1 == pos {
			cycleNode = curr
		}
	}
	if pos >= 0 {
		curr.Next = cycleNode
	}
	return head
}

func main() {
	list1 := buildList([]int{3, 2, 0, -4}, 1)
	fmt.Println(hasCycle(list1)) // true

	list2 := buildList([]int{1, 2}, 0)
	fmt.Println(hasCycle(list2)) // true

	list3 := buildList([]int{1}, -1)
	fmt.Println(hasCycle(list3)) // false
}
