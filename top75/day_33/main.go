package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Maximum Depth of Binary Tree
// Technique: DFS recursion
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(maxDepth(root)) // 3

	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println(maxDepth(root2)) // 2
}
