package main

import (
	"fmt"
	"sort"
)

/*
Problem: Contains Duplicate

Given an integer array nums,
return true if any value appears at least twice in the array,
and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:
- 1 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
*/

/*
Brute Force Solution
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
func containsDuplicateBrute(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

/*
Hash Map Solution (Counting Occurrences)
Time Complexity: O(n)
Space Complexity: O(n)
*/
func containsDuplicateMap(nums []int) bool {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
		if counts[num] > 1 {
			return true
		}
	}
	return false
}

/*
Hash Set Solution (Best Practice)
Time Complexity: O(n)
Space Complexity: O(n)
*/
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

/*
Sorting Solution
Time Complexity: O(n log n)
Space Complexity: O(1) or O(n) depending on sort implementation
*/
func containsDuplicateSort(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 2, 4}

	fmt.Println(containsDuplicateBrute(nums1)) // true
	fmt.Println(containsDuplicateMap(nums1))   // true
	fmt.Println(containsDuplicate(nums1))      // true
	fmt.Println(containsDuplicateSort(nums1))  // true

	fmt.Println(containsDuplicateBrute(nums2)) // false
	fmt.Println(containsDuplicateMap(nums2))   // false
	fmt.Println(containsDuplicate(nums2))      // false
	fmt.Println(containsDuplicateSort(nums2))  // false

	fmt.Println(containsDuplicateBrute(nums3)) // true
	fmt.Println(containsDuplicateMap(nums3))   // true
	fmt.Println(containsDuplicate(nums3))      // true
	fmt.Println(containsDuplicateSort(nums3))  // true
}
