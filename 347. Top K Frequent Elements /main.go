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
	result := make([]int, 0)

	for _, num := range nums {
		frequentMap[num]++
	}

	fmt.Println(frequentMap)

	for i := 1; i <= k; i++ {
		fmt.Println(i)
		fmt.Println(frequentMap[frequentMap[i]])
		result = append(result, frequentMap[frequentMap[i]])
	}

	return result
}

func main() {
	// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1,2]
	// fmt.Println(topKFrequent([]int{1}, 1))                // [1]
	fmt.Println(topKFrequent([]int{-1, -1}, 1)) // [-1]

}
