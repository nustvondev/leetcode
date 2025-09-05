package main

import "fmt"

/*
Problem: 238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal 
to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

Constraints:
- 2 <= nums.length <= 10^5
- -30 <= nums[i] <= 30
- Answer[i] guaranteed to fit in 32-bit integer.
- Must run in O(n), without division.
*/

/*
Brute Force (O(n^2))
Idea:
- For each index i, multiply all other elements except nums[i].
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
func productExceptSelfV1(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		product := 1
		for j := 0; j < n; j++ {
			if i != j {
				product *= nums[j]
			}
		}
		res[i] = product
	}
	return res
}

/*
Prefix & Suffix Arrays (O(n), O(n) space)
Idea:
- Precompute prefix[i] = product of nums[0..i-1]
- Precompute suffix[i] = product of nums[i+1..n-1]
- Result[i] = prefix[i] * suffix[i]
Time Complexity: O(n)
Space Complexity: O(n)
*/
func productExceptSelfV2(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)
	res := make([]int, n)

	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		res[i] = prefix[i] * suffix[i]
	}
	return res
}

/*
Optimized Prefix & Suffix (O(n), O(1) extra space)
Idea:
- Use result array as prefix storage.
- Then multiply suffix on the fly in a second pass.
Time Complexity: O(n)
Space Complexity: O(1) (excluding result array)
*/
func productExceptSelfV3(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// prefix products
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{-1, 1, 0, -3, 3}

	fmt.Println("Brute Force :", productExceptSelfV1(nums1)) // [24,12,8,6]
	fmt.Println("Prefix/Suff :", productExceptSelfV2(nums1)) // [24,12,8,6]
	fmt.Println("Optimized   :", productExceptSelfV3(nums1)) // [24,12,8,6]

	fmt.Println("Optimized 2 :", productExceptSelfV3(nums2)) // [0,0,9,0,0]
}
