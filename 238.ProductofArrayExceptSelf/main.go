package main

import "fmt"

//Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
//The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
//You must write an algorithm that runs in O(n) time and without using the division operation.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3,4]
//Output: [24,12,8,6]
//
//Example 2:
//
//Input: nums = [-1,1,0,-3,3]
//Output: [0,0,9,0,0]

func productExceptSelf(nums []int) []int {
	result1 := make([]int, len(nums))
	// result2 := make([]int, len(nums))
	// result3 := []int{}

	prefix := 1
	postfix := 1

	for i := 0; i < len(result1); i++ {
		result1[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result1[i] *= postfix
		postfix *= nums[i]
	}

	// for i := len(nums) - 1; i >= 0; i-- {
	// 	result2[i] = postfix
	// 	postfix *= nums[i]
	// }

	// for i := 0; i < len(nums); i++ {
	// 	res := result1[i] * result2[i]
	// 	result3 = append(result3, res)
	// }

	return result1
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) //[24,12,8,6]
}
