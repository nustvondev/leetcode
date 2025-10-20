package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec struct for serialization/deserialization
type Codec struct{}

// Serialize the tree to a single string (preorder traversal)
func (c *Codec) serialize(root *TreeNode) string {
	var sb []string
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb = append(sb, "null")
			return
		}
		sb = append(sb, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(sb, ",")
}

// Deserialize converts string back to tree
func (c *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	var i int
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if values[i] == "null" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(values[i])
		i++
		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}

// Helper function for testing
func printPreorder(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printPreorder(root.Left)
	printPreorder(root.Right)
}

func main() {
	codec := Codec{}
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
	}

	data := codec.serialize(root)
	fmt.Println("Serialized:", data)

	node := codec.deserialize(data)
	fmt.Print("Deserialized (Preorder): ")
	printPreorder(node)
	fmt.Println()
}
