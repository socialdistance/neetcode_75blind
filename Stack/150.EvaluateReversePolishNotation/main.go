package main

import (
	"fmt"
	"strconv"
)

// You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

// Evaluate the expression. Return an integer that represents the value of the expression.

// Note that:

//     The valid operators are '+', '-', '*', and '/'.
//     Each operand may be an integer or another expression.
//     The division between two integers always truncates toward zero.
//     There will not be any division by zero.
//     The input represents a valid arithmetic expression in a reverse polish notation.
//     The answer and all the intermediate calculations can be represented in a 32-bit integer.

// Example 1:

// Input: tokens = ["2","1","+","3","*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9

// Example 2:

// Input: tokens = ["4","13","5","/","+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6

// Example 3:

// Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// Output: 22
// Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	var a, b int

	for _, c := range tokens {
		switch c {
		case "+":
			a, b, stack = getElementsAndPop(stack)
			stack = append(stack, a+b)
		case "-":
			a, b, stack = getElementsAndPop(stack)
			stack = append(stack, a-b)
		case "*":
			a, b, stack = getElementsAndPop(stack)
			stack = append(stack, a*b)
		case "/":
			a, b, stack = getElementsAndPop(stack)
			stack = append(stack, a/b)
		default:
			i, _ := strconv.Atoi(c)
			stack = append(stack, i)
		}
	}

	return stack[0]
}

func getElementsAndPop(stack []int) (int, int, []int) {
	a := stack[len(stack)-2]
	b := stack[len(stack)-1]
	stack = stack[:len(stack)-2]

	return a, b, stack
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))  // 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"})) // 6
}
