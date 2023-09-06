package main

import "fmt"

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
	top   int
	chars []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		chars: make([]int, 0),
		top:   0,
		min:   0,
	}

}

func (this *MinStack) Push(val int) {
	this.top++

	this.chars = append(this.chars, val)

	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {

	this.chars = this.chars[:len(this.chars)-1]
	this.top--
	if this.chars[len(this.chars)-1] < this.min {
		this.min = this.chars[len(this.chars)-1]
	}
}

func (this *MinStack) Top() int {
	return this.chars[this.top-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	minStack := new(MinStack)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
