package main

import (
	"fmt"
)

//Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//Example 1:
//
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//
//Example 2:
//
//Input: nums = [1], k = 1
//Output: [1]

func topKFrequent(nums []int, k int) []int {
	frequentMap := make(map[int]int, 0)
	result := make([]int, k)

	for _, num := range nums {
		frequentMap[num]++
	}

	fmt.Println(frequentMap)

	for k, v := range frequentMap {
		fmt.Println(k, v)
	}

	fmt.Println(result)

	return []int{}
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1,2]
}
