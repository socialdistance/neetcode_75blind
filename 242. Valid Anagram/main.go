package main

import "fmt"

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:

// Input: s = "rat", t = "car"
// Output: false

func isAnagram(s string, t string) bool {
	// var freq [26]int

	// for idx := 0; idx < len(s); idx++ {
	// 	freq[s[idx]-'a']++
	// 	freq[t[idx]-'a']--
	// }

	// for i := 0; i < len(freq); i++ {
	// 	if freq[i] != 0 {
	// 		return false
	// 	}
	// }

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("aacc", "ccac"))
}
