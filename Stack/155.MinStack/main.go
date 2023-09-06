package main

import (
	"fmt"
	"math"
)

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

//     MinStack() initializes the stack object.
//     void push(int val) pushes the element val onto the stack.
//     void pop() removes the element on the top of the stack.
//     int top() gets the top element of the stack.
//     int getMin() retrieves the minimum element in the stack.

// You must implement a solution with O(1) time complexity for each function.

// Example 1:

// Input
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// Output
// [null,null,null,null,-3,null,0,-2]

// Explanation
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin(); // return -3
// minStack.pop();
// minStack.top();    // return 0
// minStack.getMin(); // return -2

type MinStack struct {
	top     int
	chars   []int
	min     int
	lastMin []int
}

func Constructor() MinStack {
	return MinStack{
		chars:   make([]int, 0),
		top:     -1,
		min:     math.MaxUint32,
		lastMin: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.top++

	this.chars = append(this.chars, val)

	if val < this.min {
		this.min = val
		this.lastMin = append(this.lastMin, this.min)
	}

}

func (this *MinStack) Pop() {
	if this.chars[len(this.chars)-1] == this.min {
		this.min = this.lastMin[len(this.lastMin)-1]
		this.lastMin = this.lastMin[:len(this.lastMin)-1]
	}

	this.chars = this.chars[:len(this.chars)-1]
	this.top--
}

func (this *MinStack) Top() int {
	return this.chars[this.top]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2

	// minStack := Constructor()
	// minStack.Push(2)
	// minStack.Push(0)
	// minStack.Push(3)
	// minStack.Push(0)
	// fmt.Println(minStack.GetMin()) // return 0
	// minStack.Pop()
	// fmt.Println(minStack.GetMin()) // return 0
	// minStack.Pop()
	// fmt.Println(minStack.GetMin()) // return 0
	// minStack.Pop()
	// fmt.Println(minStack.GetMin()) // return 2
}
