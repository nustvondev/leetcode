package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Subtree of Another Tree
// Technique: DFS + Tree Comparison
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// helper: compare two trees
func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}

	subRoot := &TreeNode{
		Val: 4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	fmt.Println(isSubtree(root, subRoot)) // true
}
