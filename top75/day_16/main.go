package main

import "fmt"

/*
Binary Search – O(log n)
- The array is sorted but rotated.
- We want to find the minimum element.
- Key insight: one half is always sorted.
- If nums[mid] > nums[right] -> min is in the right half.
- Else -> min is in the left half (including mid).
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// Minimum must be in the right half
			left = mid + 1
		} else {
			// Minimum is in the left half (including mid)
			right = mid
		}
	}
	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // 11
}
