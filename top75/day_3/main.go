package main

import "fmt"

/*
Problem: Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock 
and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. 
If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

Constraints:
- 1 <= prices.length <= 10^5
- 0 <= prices[i] <= 10^4
*/

/*
Brute Force Solution
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
func maxProfitV1(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

/*
Optimal One-Pass Solution
Time Complexity: O(n)
Space Complexity: O(1)
*/
func maxProfitV2(prices []int) int {
	minPrice := int(^uint(0) >> 1) // max int
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfitV1(prices1)) // Output: 5
	fmt.Println(maxProfitV2(prices1)) // Output: 5

	fmt.Println(maxProfitV1(prices2)) // Output: 0
	fmt.Println(maxProfitV2(prices2)) // Output: 0
}
