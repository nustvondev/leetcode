package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Invert Binary Tree
// Technique: DFS (Recursion)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// Hoán đổi cây con trái và phải
	root.Left, root.Right = root.Right, root.Left

	// Gọi đệ quy cho 2 cây con
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// Helper: Inorder traversal (for output)
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

func main() {
	root := &TreeNode{4,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}},
	}

	fmt.Print("Before invert: ")
	inorder(root)
	fmt.Println()

	inverted := invertTree(root)

	fmt.Print("After invert: ")
	inorder(inverted)
	fmt.Println()
}
