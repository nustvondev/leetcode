package main

import "fmt"

/*
Brute Force – O(n^2)
- Check all pairs (i, j), calculate area = min(height[i], height[j]) * (j - i).
- Too slow for n=1e5.
*/
func maxAreaBruteForce(height []int) int {
	n := len(height)
	maxArea := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			h := height[i]
			if height[j] < h {
				h = height[j]
			}
			area := h * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

/*
Two Pointers – O(n)
Idea:
- Start with left=0, right=n-1.
- Compute area, update max.
- Move the pointer with smaller height inward (because it limits container).
- Continue until left >= right.
*/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
}
