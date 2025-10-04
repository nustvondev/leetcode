package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}

	var dfs func(int,int,[][]bool)
	dfs = func(r, c int, visited [][]bool) {
		visited[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			// chỉ đi đến ô cao hơn hoặc bằng (để đảm bảo nước chảy ngược hợp lệ)
			if visited[nr][nc] || heights[nr][nc] < heights[r][c] {
				continue
			}
			dfs(nr, nc, visited)
		}
	}

	// Pacific (top row + left col)
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pacific)
	}

	// Atlantic (bottom row + right col)
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, atlantic)
	}

	// intersection
	res := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func main() {
	heights := [][]int{
		{1,2,2,3,5},
		{3,2,3,4,4},
		{2,4,5,3,1},
		{6,7,1,4,5},
		{5,1,1,2,4},
	}
	fmt.Println(pacificAtlantic(heights))
}
