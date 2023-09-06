package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

//     Open brackets must be closed by the same type of brackets.
//     Open brackets must be closed in the correct order.
//     Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"
// Output: true

// Example 2:

// Input: s = "()[]{}"
// Output: true

// Example 3:

// Input: s = "(]"
// Output: false

func isValid(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		pair, ok := pairs[char]
		if !ok {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != pair {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}
