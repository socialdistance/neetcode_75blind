package main

import "fmt"

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

func maxArea(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)

		if area > res {
			res = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
