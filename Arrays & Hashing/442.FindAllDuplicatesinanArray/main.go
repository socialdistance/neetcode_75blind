package main

import (
	"math"
)

// Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.

// You must write an algorithm that runs in O(n) time and uses only constant extra space.

// Example 1:

// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [2,3]

// Example 2:

// Input: nums = [1,1,2]
// Output: [1]

// Example 3:

// Input: nums = [1]
// Output: []

func findDuplicates(nums []int) []int {
	result := make([]int, 0)

	for _, num := range nums {
		value := math.Abs(float64(num))
		index := value - 1

		if nums[int(index)] < 0 {
			result = append(result, int(value))
		}

		nums[int(index)] *= -1
	}

	return result
}

// func findDuplicates(nums []int) []int {
// 	resultMap := make(map[int]struct{})
// 	resultSlice := make([]int, 0)

// 	for _, num := range nums {
// 		if _, ok := resultMap[num]; ok {
// 			resultSlice = append(resultSlice, num)
// 		}

// 		resultMap[num] = struct{}{}
// 	}

// 	return resultSlice
// }

func main() {
	// fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})) // [2, 3]
	// fmt.Println(findDuplicates([]int{1, 1, 2}))                // [1]
	// fmt.Println(findDuplicates([]int{1}))                      // []
}
