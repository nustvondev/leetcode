package main

import "fmt"
import "math"

// Validate Binary Search Tree
// Technique: DFS with value bounds
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}
	val := int64(node.Val)
	if val <= min || val >= max {
		return false
	}
	return validate(node.Left, min, val) && validate(node.Right, val, max)
}

func main() {
	root1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(isValidBST(root1)) // true

	root2 := &TreeNode{5,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}},
	}
	fmt.Println(isValidBST(root2)) // false
}
