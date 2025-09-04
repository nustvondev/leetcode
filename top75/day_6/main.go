package main

import "fmt"

/*
Problem: Maximum Subarray

Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23

Constraints:
- 1 <= nums.length <= 10^5
- -10^4 <= nums[i] <= 10^4

Follow up:
If you have figured out the O(n) solution, try coding another solution using the
divide and conquer approach, which is more subtle.
*/

/*
Brute Force Solution (O(n^3))
Idea:
- Check all subarrays, compute their sum, and track the maximum.
Time Complexity: O(n^3)
Space Complexity: O(1)
*/
func maxSubArrayV1(nums []int) int {
	n := len(nums)
	maxSum := nums[0]

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

/*
Improved Brute Force with Prefix Sum (O(n^2))
Idea:
- Use prefix sums to calculate subarray sums quickly.
- sum(i..j) = prefix[j+1] - prefix[i]
Time Complexity: O(n^2)
Space Complexity: O(n)
*/
func maxSubArrayV2(nums []int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	maxSum := nums[0]
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := prefix[j+1] - prefix[i]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

/*
Kadane’s Algorithm (Best Practice O(n))
Idea:
- For each element, decide whether to start a new subarray
  or extend the previous subarray.
- Track the global maximum.
Time Complexity: O(n)
Space Complexity: O(1)
*/
func maxSubArrayV3(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currSum+nums[i] > nums[i] {
			currSum = currSum + nums[i]
		} else {
			currSum = nums[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}

/*
Divide and Conquer (O(n log n))
Idea:
- Divide array into two halves.
- Maximum subarray is either:
  1. In left half
  2. In right half
  3. Crossing the middle
Time Complexity: O(n log n)
Space Complexity: O(log n) due to recursion
*/
func maxCrossingSum(nums []int, left, mid, right int) int {
	leftSum := -1 << 31
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	rightSum := -1 << 31
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return leftSum + rightSum
}

func maxSubArrayHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	leftMax := maxSubArrayHelper(nums, left, mid)
	rightMax := maxSubArrayHelper(nums, mid+1, right)
	crossMax := maxCrossingSum(nums, left, mid, right)

	if leftMax >= rightMax && leftMax >= crossMax {
		return leftMax
	} else if rightMax >= leftMax && rightMax >= crossMax {
		return rightMax
	} else {
		return crossMax
	}
}

func maxSubArrayV4(nums []int) int {
	return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func main() {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums2 := []int{1}
	nums3 := []int{5, 4, -1, 7, 8}

	fmt.Println("Brute Force:", maxSubArrayV1(nums1)) // 6
	fmt.Println("Prefix Sum :", maxSubArrayV2(nums1)) // 6
	fmt.Println("Kadane     :", maxSubArrayV3(nums1)) // 6
	fmt.Println("Divide&Conq:", maxSubArrayV4(nums1)) // 6

	fmt.Println("Kadane [1]:", maxSubArrayV3(nums2))  // 1
	fmt.Println("Kadane [2]:", maxSubArrayV3(nums3))  // 23
}
