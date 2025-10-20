package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Binary Tree Maximum Path Sum
// Technique: DFS + Recursion
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		currentPath := node.Val + left + right
		maxSum = max(maxSum, currentPath)
		return node.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(maxPathSum(root)) // 6

	root2 := &TreeNode{
		Val: -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(maxPathSum(root2)) // 42
}
