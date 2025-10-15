package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// Clone Graph - DFS Approach
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if clone, ok := visited[node]; ok {
		return clone
	}

	clone := &Node{Val: node.Val}
	visited[node] = clone

	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(neighbor, visited))
	}
	return clone
}

// Helper for demonstration
func printGraph(node *Node, visited map[*Node]bool) {
	if node == nil || visited[node] {
		return
	}
	visited[node] = true
	fmt.Printf("Node %d -> ", node.Val)
	for _, n := range node.Neighbors {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println()
	for _, n := range node.Neighbors {
		printGraph(n, visited)
	}
}

func main() {
	// Example graph: 1--2, 2--3, 3--4, 4--1 (square)
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	clone := cloneGraph(n1)
	fmt.Println("Original Graph:")
	printGraph(n1, make(map[*Node]bool))

	fmt.Println("Cloned Graph:")
	printGraph(clone, make(map[*Node]bool))
}
