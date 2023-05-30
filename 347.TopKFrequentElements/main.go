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
	countMap := make(map[int]int, 0)
	resultSlice := make([][]int, len(nums)+1)
	var result []int

	for _, value := range nums {
		if _, ok := countMap[value]; !ok {
			countMap[value] = 1
		} else {
			countMap[value]++
		}
	}

	for count, value := range countMap {
		resultSlice[value] = append(resultSlice[value], count)
	}

	for i := len(resultSlice) - 1; i > 0; i-- {
		result = append(result, resultSlice[i]...)
		if k == len(result) {
			return result
		}
	}

	return []int{}
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1,2]
	fmt.Println(topKFrequent([]int{1}, 1))                // [1]
	fmt.Println(topKFrequent([]int{-1, -1}, 1))           // [-1]
}
