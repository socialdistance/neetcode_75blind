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

func topKFrequent(nums []int, k int) (res []int) {
	frequentMap := map[int]int{}
	resultMap := make([][]int, len(nums)+1)

	for _, num := range nums {
		// if count, ok := frequentMap[num]; ok {
		// 	frequentMap[num] = count + 1
		// } else {
		// 	frequentMap[num] = 1
		// }
		frequentMap[num]++
	}

	for count, value := range frequentMap {
		resultMap[value] = append(resultMap[value], count)
	}

	for i := len(resultMap) - 1; i > 0; i-- {
		res = append(res, resultMap[i]...)
		if k == len(res) {
			return res
		}
	}

	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1,2]
	// fmt.Println(topKFrequent([]int{1}, 1))                // [1]
	fmt.Println(topKFrequent([]int{-1, -1}, 1)) // [-1]
}
