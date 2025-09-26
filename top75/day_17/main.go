package main

import "fmt"

/*
Sliding Window + Frequency Count
- We want the longest substring where we can replace <= k chars to make all same.
- For each window [l..r], track the count of the most frequent char (maxCount).
- Condition: (window size - maxCount) <= k
  -> If not, shrink the window from the left.
*/
func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left, maxCount, result := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		if count[s[right]] > maxCount {
			maxCount = count[s[right]]
		}

		// Check if window is invalid
		if (right-left+1)-maxCount > k {
			count[s[left]]--
			left++
		}

		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))    // 4
	fmt.Println(characterReplacement("AABABBA", 1)) // 4
}
