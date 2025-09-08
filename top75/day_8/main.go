package main

import (
	"fmt"
	"sort"
)

/*
Problem: 15. 3Sum

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] 
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

The solution set must not contain duplicate triplets.

Constraints:
- 3 <= nums.length <= 3000
- -10^5 <= nums[i] <= 10^5
*/

/*
Brute Force (O(n^3))
Idea:
- Iterate over all triplets (i, j, k).
- If sum == 0, add to result.
- Use a set to avoid duplicates.
Time Complexity: O(n^3)
Space Complexity: O(n^2) for storing triplets.
*/
func threeSumV1(nums []int) [][]int {
	n := len(nums)
	set := make(map[[3]int]bool)
	var res [][]int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					sort.Ints(triplet)
					key := [3]int{triplet[0], triplet[1], triplet[2]}
					if !set[key] {
						set[key] = true
						res = append(res, triplet)
					}
				}
			}
		}
	}
	return res
}

/*
Sorting + Two Pointers (O(n^2)) – Best Practice
Idea:
- Sort array.
- Fix one number nums[i], then find two-sum in remaining array.
- Skip duplicates to avoid repeating triplets.
Time Complexity: O(n^2)
Space Complexity: O(1) excluding output.
*/
func threeSumV2(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate first element
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// skip duplicates
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{0, 1, 1}
	nums3 := []int{0, 0, 0}

	fmt.Println("Brute Force:", threeSumV1(nums1)) // [[-1,-1,2],[-1,0,1]]
	fmt.Println("Two Pointer:", threeSumV2(nums1)) // [[-1,-1,2],[-1,0,1]]

	fmt.Println("Two Pointer:", threeSumV2(nums2)) // []
	fmt.Println("Two Pointer:", threeSumV2(nums3)) // [[0,0,0]]
}
