package main

import "fmt"

/*
Problem: Two Sum

Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:
- 2 <= nums.length <= 10^4
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9
- Only one valid answer exists.
*/

/*
Brute Force Solution
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
func twoSumV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
Hash Map Solution (Best Practice)
Time Complexity: O(n)
Space Complexity: O(n)
*/
func twoSumV2(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, found := hashMap[complement]; found {
			return []int{idx, i}
		}
		hashMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSumV1(nums, target)) // Output: [0 1]
	fmt.Println(twoSumV2(nums, target)) // Output: [0 1]
}
