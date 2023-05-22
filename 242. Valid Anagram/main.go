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
	mapResult := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for idx := 0; idx < len(s); idx++ {
		mapResult[s[idx]]++
		mapResult[t[idx]]--
	}

	for _, v := range mapResult {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("a", "ab"))            // false
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println(isAnagram("aacc", "ccac"))       // false
}
