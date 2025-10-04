package main

import "fmt"

// Count palindromic substrings
func countSubstrings(s string) int {
	n := len(s)
	count := 0

	// helper: expand from middle
	expand := func(left, right int) {
		for left >= 0 && right < n && s[left] == s[right] {
			count++
			left--
			right++
		}
	}

	// odd length (center at i), even length (center at i, i+1)
	for i := 0; i < n; i++ {
		expand(i, i)     // odd
		expand(i, i+1)   // even
	}

	return count
}

func main() {
	fmt.Println(countSubstrings("abc")) // 3
	fmt.Println(countSubstrings("aaa")) // 6
}
