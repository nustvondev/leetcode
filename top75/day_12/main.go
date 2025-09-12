package main

import "fmt"

/*
Problem: 33. Search in Rotated Sorted Array
*/

/*
Brute Force – O(n)
Simply iterate and check.
*/
func searchBruteForce(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

/*
Optimized Binary Search – O(log n)
Idea:
- Array is rotated, so one side (left or right) is still sorted.
- We can determine which side is sorted and decide to move left/right accordingly.
Steps:
1. While left <= right:
   - mid = (left+right)/2
   - If nums[mid] == target → return mid
   - If left half sorted (nums[left] <= nums[mid]):
        * If target in [nums[left], nums[mid]] → right = mid-1
        * Else → left = mid+1
   - Else right half sorted:
        * If target in [nums[mid], nums[right]] → left = mid+1
        * Else → right = mid-1
2. Return -1 if not found.
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // left half sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right half sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(search([]int{1}, 0))                   // -1
}
