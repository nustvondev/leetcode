package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	countT := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}

	have, need := 0, len(countT)
	res := [2]int{-1, -1}
	resLen := 1<<31 - 1
	l := 0

	for r := 0; r < len(s); r++ {
		c := s[r]
		window[c]++
		if countT[c] > 0 && window[c] == countT[c] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				res = [2]int{l, r}
				resLen = r - l + 1
			}
			window[s[l]]--
			if countT[s[l]] > 0 && window[s[l]] < countT[s[l]] {
				have--
			}
			l++
		}
	}

	if resLen == 1<<31-1 {
		return ""
	}
	return s[res[0] : res[1]+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // "BANC"
	fmt.Println(minWindow("a", "a"))               // "a"
	fmt.Println(minWindow("a", "aa"))              // ""
}
