package main

import "fmt"

// Course Schedule - Topological Sort using DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	for _, p := range prerequisites {
		a, b := p[0], p[1]
		graph[b] = append(graph[b], a)
	}

	visited := make([]int, numCourses) // 0 = unvisited, 1 = visiting, 2 = done
	var dfs func(int) bool
	dfs = func(course int) bool {
		if visited[course] == 1 { // found a cycle
			return false
		}
		if visited[course] == 2 { // already checked
			return true
		}
		visited[course] = 1
		for _, next := range graph[course] {
			if !dfs(next) {
				return false
			}
		}
		visited[course] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))       // true
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}})) // false
	fmt.Println(canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}})) // true
}
