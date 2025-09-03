package main

import "fmt"

/*
Problem: Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:
- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
- 1 <= s.length <= 10^4
- s consists of parentheses only '()[]{}'.
*/

/*
Brute Force Solution (String Replace)
Idea: Keep replacing "()", "{}", "[]" with "" until no more replacements can be made.
If the final string is empty, return true; otherwise false.
Time Complexity: O(n^2) in worst case
Space Complexity: O(n)
*/
func isValidV1(s string) bool {
	prev := ""
	for s != prev {
		prev = s
		s = replacePairs(s)
	}
	return len(s) == 0
}

func replacePairs(s string) string {
	s = replaceAll(s, "()", "")
	s = replaceAll(s, "[]", "")
	s = replaceAll(s, "{}", "")
	return s
}

func replaceAll(s, old, new string) string {
	result := ""
	for i := 0; i < len(s); {
		if i+len(old) <= len(s) && s[i:i+len(old)] == old {
			result += new
			i += len(old)
		} else {
			result += string(s[i])
			i++
		}
	}
	return result
}

/*
Stack Solution (Best Practice)
Idea: Use a stack to track open brackets.
- Push expected closing bracket when encountering an opening bracket.
- When encountering a closing bracket, check stack top.
Time Complexity: O(n)
Space Complexity: O(n)
*/
func isValidV2(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		// If it's closing bracket
		if open, ok := brackets[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else {
			// Push opening bracket
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

/*
Optimized Stack Solution
Idea: Instead of pushing opening bracket, push the expected closing bracket.
This way, when encountering a closing bracket, just check directly.
Time Complexity: O(n)
Space Complexity: O(n)
*/
func isValidV3(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		}
	}

	return len(stack) == 0
}

func main() {
	// Example test cases
	fmt.Println(isValidV1("()"))      // true
	fmt.Println(isValidV1("()[]{}"))  // true
	fmt.Println(isValidV1("(]"))      // false
	fmt.Println(isValidV1("([])"))    // true
	fmt.Println(isValidV1("([)]"))    // false

	fmt.Println(isValidV2("()"))      // true
	fmt.Println(isValidV2("()[]{}"))  // true
	fmt.Println(isValidV2("(]"))      // false
	fmt.Println(isValidV2("([])"))    // true
	fmt.Println(isValidV2("([)]"))    // false

	fmt.Println(isValidV3("()"))      // true
	fmt.Println(isValidV3("()[]{}"))  // true
	fmt.Println(isValidV3("(]"))      // false
	fmt.Println(isValidV3("([])"))    // true
	fmt.Println(isValidV3("([)]"))    // false
}
