package main

import "fmt"

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:

// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.

// Example 2:

// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achive this answer too.

// func characterReplacement(s string, k int) int {
// 	count := make(map[string]int)
// 	res := 0

// 	left := 0
// 	maxFreq := 0

// 	for r, _ := range s {
// 		count[string(s[r])] = 1 + count[string(s[r])]
// 		maxFreq = max(maxFreq, count[string(s[r])])

// 		if (r-left+1)-maxFreq > k {
// 			count[string(s[r])] -= 1
// 			left++
// 		}

// 		res = max(res, r-left+1)
// 	}

// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res := 0

	l := 0
	maxf := 0
	for r, _ := range s {
		count[s[r]] = 1 + count[s[r]]
		maxf = max(maxf, count[s[r]])

		if (r-l+1)-maxf > k {
			count[s[l]] -= 1
			l++
		}
		res = max(res, r-l+1)
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
	// fmt.Println(characterReplacement("ABAB", 2))    // 4 Replace the two 'A's with two 'B's or vice versa.
	fmt.Println(characterReplacement("AABABBA", 1)) // 4

}
