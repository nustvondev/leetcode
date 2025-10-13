package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Construct Binary Tree from Preorder and Inorder Traversal
// Recursive divide & conquer approach using hashmap for fast lookup
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	preIdx := 0
	var helper func(int, int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		rootVal := preorder[preIdx]
		preIdx++
		root := &TreeNode{Val: rootVal}
		idx := indexMap[rootVal]
		root.Left = helper(left, idx-1)
		root.Right = helper(idx+1, right)
		return root
	}

	return helper(0, len(inorder)-1)
}

// Inorder traversal to verify result
func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Print(root.Val, " ")
	printInorder(root.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	printInorder(root) // Output: 9 3 15 20 7
}
