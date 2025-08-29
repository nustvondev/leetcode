package main

import "fmt"

/*
Problem: Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
- 1 <= s.length, t.length <= 5 * 10^4
- s and t consist of lowercase English letters.

Follow up: What if the inputs contain Unicode characters? 
How would you adapt your solution to such a case?
*/

/*
Brute Force Solution
Idea: Sort both strings and compare
Time Complexity: O(n log n)
Space Complexity: O(1) or O(n) depending on sorting implementation
*/
func isAnagramV1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// Convert to slice of bytes (Go strings are immutable)
	sBytes := []byte(s)
	tBytes := []byte(t)

	// Sort both slices
	for i := 0; i < len(sBytes); i++ {
		for j := i + 1; j < len(sBytes); j++ {
			if sBytes[i] > sBytes[j] {
				sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
			}
			if tBytes[i] > tBytes[j] {
				tBytes[i], tBytes[j] = tBytes[j], tBytes[i]
			}
		}
	}
	// Compare
	for i := range sBytes {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}
	return true
}

/*
Hash Map Solution (Best Practice)
Idea: Count frequency of characters in s and t, then compare
Time Complexity: O(n)
Space Complexity: O(1) because only lowercase English letters (max 26)
*/
func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

/*
Unicode-Friendly Solution
Idea: Use a hash map of runes to count frequency (works for any Unicode character)
Time Complexity: O(n)
Space Complexity: O(k) where k is the number of unique characters
*/
func isAnagramV3(s string, t string) bool {
	if len([]rune(s)) != len([]rune(t)) {
		return false
	}

	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

func main() {
	// Example test cases
	fmt.Println(isAnagramV1("anagram", "nagaram")) // true
	fmt.Println(isAnagramV1("rat", "car"))         // false

	fmt.Println(isAnagramV2("anagram", "nagaram")) // true
	fmt.Println(isAnagramV2("rat", "car"))         // false

	fmt.Println(isAnagramV3("你好世界", "世界你好")) // true (Unicode case)
}
