package main

import "fmt"

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
// Example 1:

// Input: nums = [1,2,3,1]
// Output: true

// Example 2:

// Input: nums = [1,2,3,4]
// Output: false

// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

func containsDuplicate(nums []int) bool {
	numsDuplicate := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := numsDuplicate[v]; ok {
			return true
		} else {
			numsDuplicate[v] = struct{}{}
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

}
