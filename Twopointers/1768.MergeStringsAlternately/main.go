package main

import (
	"fmt"
	"strings"
)

// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

// Example 1:

// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r

// Example 2:

// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b
// word2:    p   q   r   s
// merged: a p b q   r   s

// Example 3:

// Input: word1 = "abcd", word2 = "pq"
// Output: "apbqcd"
// Explanation: Notice that as word1 is longer, "cd" is appended to the end.
// word1:  a   b   c   d
// word2:    p   q
// merged: a p b q c   d

func mergeAlternately(word1 string, word2 string) string {
	left := 0
	right := 0
	var result []string

	if len(word1) <= len(word2) {
		for left < len(word1) {
			result = append(result, string(word1[left]), string(word2[right]))
			left++
			right++
		}

		result = append(result, word2[right:])
	}

	fmt.Println(result)

	return strings.Join(result, "")
}

func main() {
	// fmt.Println(mergeAlternately("abc", "pqr")) // apbqcr
	// fmt.Println(mergeAlternately("ab", "pqrs")) // apbqrs
	fmt.Println(mergeAlternately("abcd", "pq")) // apbqcd
}
