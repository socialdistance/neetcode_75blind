package main

import "fmt"

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//
// You must implement a solution with a linear runtime complexity and use only constant extra space.
//
//
//
// Example 1:
//
// Input: nums = [2,2,1]
//
// Output: 1
//
// Example 2:
//
// Input: nums = [4,1,2,1,2]
//
// Output: 4
//
// Example 3:
//
// Input: nums = [1]
//
// Output: 1

func singleNumber(nums []int) int {
	cnt := map[int]int{}

	for _, num := range nums {
		cnt[num]++
	}

	for key, val := range cnt {
		if val == 1 {
			return key
		}
	}

	return 0
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))       // 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2})) // 4
	fmt.Println(singleNumber([]int{1}))             // 1
	fmt.Println(singleNumber([]int{-1}))            // -1
}
