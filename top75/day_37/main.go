package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Lowest Common Ancestor of a Binary Search Tree
// Technique: Use BST property to navigate
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// If both nodes are smaller than root, go left
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	// If both nodes are greater than root, go right
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// Otherwise, this node is the split point → LCA
	return root
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 3},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}

	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 8}

	lca := lowestCommonAncestor(root, p, q)
	fmt.Println(lca.Val) // Output: 6
}
