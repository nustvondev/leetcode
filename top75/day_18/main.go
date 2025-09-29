package main

import "fmt"

/*
Sliding Window + Hash Map
- Dùng hai pointer [left..right] duyệt qua chuỗi.
- Map lưu vị trí cuối cùng của mỗi ký tự.
- Nếu gặp ký tự trùng trong window, dịch left sang bên phải.
*/
func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[byte]int)
	left, result := 0, 0

	for right := 0; right < len(s); right++ {
		if idx, ok := lastIndex[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		lastIndex[s[right]] = right
		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}
