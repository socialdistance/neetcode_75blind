package main

import (
	"fmt"
	"unicode"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// func isPalindrome(s string) bool {
// 	var result, resultReverse strings.Builder
// 	result.Grow(len(s))

// 	lowerResult := strings.ToLower(s)

// 	for i := 0; i < len(lowerResult); i++ {
// 		b := lowerResult[i]
// 		if ('a' <= b && b <= 'z') ||
// 			('A' <= b && b <= 'Z') ||
// 			('0' <= b && b <= '9') {
// 			result.WriteByte(b)
// 		}
// 	}

// 	resultReverse.Grow(len(result.String()))

// 	for i := len(result.String()) - 1; i >= 0; i-- {
// 		b := result.String()[i]
// 		resultReverse.WriteByte(b)
// 	}

// 	if result.String() == resultReverse.String() {
// 		return true
// 	}

// 	return false
// }

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	result := []rune(s)

	for i < j {
		left := unicode.ToLower(result[i])
		right := unicode.ToLower(result[j])

		if !isLetterOrDigit(left) {
			i++
			continue
		}

		if !isLetterOrDigit(right) {
			j--
			continue
		}

		if left != right {
			return false
		}

		i++
		j--
	}

	return true
}

func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("raceacar"))                       // false
	// fmt.Println(isPalindrome(" "))                              // true

}
