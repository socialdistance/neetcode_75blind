package main

import "fmt"

// Given a string s, return true if the s can be palindrome after deleting at most one character from it.

// Example 1:

// Input: s = "aba"
// Output: true

// Example 2:

// Input: s = "abca"
// Output: true
// Explanation: You could delete the character 'c'.

// Example 3:

// Input: s = "abc"
// Output: false

func validPalindrome(s string) bool {

	return false
}

func main() {
	fmt.Println(validPalindrome("aba"))  // true
	fmt.Println(validPalindrome("abca")) // true You could delete the character 'c'.

}
