package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Same Tree
// Technique: Recursion
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both empty
	if p == nil && q == nil {
		return true
	}
	// One empty, one not
	if p == nil || q == nil {
		return false
	}
	// Compare current value + left subtree + right subtree
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(isSameTree(p, q)) // true

	p2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println(isSameTree(p2, q2)) // false
}
