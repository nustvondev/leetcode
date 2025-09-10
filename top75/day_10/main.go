package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Problem: 49. Group Anagrams

Given an array of strings strs, group the anagrams together. 
You can return the answer in any order.

Constraints:
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters.
*/

/*
Brute Force (O(n^2 * k log k))
Idea:
- For each word, compare with others by sorting.
- Group if sorted strings are equal.
Not efficient for large input.
*/

/*
Optimized (HashMap with Sorted Key)
Idea:
- For each word, sort its characters to form a key.
- Use map[key] = list of words.
- Collect all values of the map.
Time Complexity: O(n * k log k)
Space Complexity: O(n * k)
*/
func groupAnagramsV1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, word := range strs {
		s := strings.Split(word, "")
		sort.Strings(s)
		key := strings.Join(s, "")
		m[key] = append(m[key], word)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

/*
Optimized (HashMap with Count Key)
Idea:
- Each word can be represented by a frequency count of 26 letters.
- Build key from counts (like "a2b1c0...").
- Use as hashmap key.
Time Complexity: O(n * k)
Space Complexity: O(n * k)
*/
func groupAnagramsV2(strs []string) [][]string {
	m := make(map[string][]string)
	for _, word := range strs {
		count := make([]int, 26)
		for _, ch := range word {
			count[ch-'a']++
		}
		// convert count to string key
		key := make([]string, 26)
		for i := 0; i < 26; i++ {
			key[i] = fmt.Sprintf("%d#", count[i])
		}
		k := strings.Join(key, "")
		m[k] = append(m[k], word)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagramsV1([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// [["eat","tea","ate"], ["tan","nat"], ["bat"]]

	fmt.Println(groupAnagramsV2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// [["eat","tea","ate"], ["tan","nat"], ["bat"]]

	fmt.Println(groupAnagramsV2([]string{""}))   // [[""]]
	fmt.Println(groupAnagramsV2([]string{"a"})) // [["a"]]
}
