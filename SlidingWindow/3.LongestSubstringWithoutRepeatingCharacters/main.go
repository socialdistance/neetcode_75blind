package main

import "fmt"

// Given a string s, find the length of the longest
// substring
// without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	left := 0
	res := 0

	for r, _ := range s {
		for charSet[s[r]] {
			delete(charSet, s[left])
			left++
		}
		charSet[s[r]] = true
		res = max(res, r-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // The answer is "abc", with the length of 3
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))   // The answer is "wke", with the length of 3
}
