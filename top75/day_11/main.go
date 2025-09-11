package main

import "fmt"

/*
Problem: 152. Maximum Product Subarray

Given an integer array nums, find a subarray that has the largest product, and return the product.
*/

/*
Brute Force (O(n^2))
Idea:
- For each subarray, compute product and track max.
- Time: O(n^2), Space: O(1).
*/
func maxProductBruteForce(nums []int) int {
	n := len(nums)
	maxProd := nums[0]

	for i := 0; i < n; i++ {
		prod := 1
		for j := i; j < n; j++ {
			prod *= nums[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	return maxProd
}

/*
Optimized DP (O(n)) – Best Practice
Idea:
- Negative numbers can flip max/min.
- Track both maxSoFar and minSoFar at each step.
- Transition:
  maxSoFar = max(nums[i], nums[i]*prevMax, nums[i]*prevMin)
  minSoFar = min(nums[i], nums[i]*prevMax, nums[i]*prevMin)
- Keep track of global max.
Time: O(n), Space: O(1).
*/
func maxProduct(nums []int) int {
	n := len(nums)
	maxSoFar, minSoFar := nums[0], nums[0]
	result := nums[0]

	for i := 1; i < n; i++ {
		cur := nums[i]
		if cur < 0 {
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}
		maxSoFar = max(cur, maxSoFar*cur)
		minSoFar = min(cur, minSoFar*cur)

		if maxSoFar > result {
			result = maxSoFar
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProductBruteForce([]int{2, 3, -2, 4})) // 6
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))           // 6

	fmt.Println(maxProduct([]int{-2, 0, -1})) // 0
}
