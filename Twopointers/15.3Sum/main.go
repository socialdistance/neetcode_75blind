package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:

// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:

// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

func threeSum(nums []int) [][]int {
	n := len(nums)

	sort.Ints(nums)

	var result [][]int

	// n-2 потому что последние значения нам не нужны в данном случае
	for i := 0; i < n-2; i++ {
		// Проверяем на дубликаты в левой части и проверяем соседей слева
		// А так же гарантируем, что счетчик пойдет со 2 элемента
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			// Если sum = 0, значит заполняем результирующий массив
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// После заполнения результирующего массива, мы должны сместить
				// правый указатель влево
				right--

				// Проверка на дубликаты справа
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else if sum > 0 {
				// Если результат > 0, смещаем указатель вправо, уменьшая счетчик
				right--
			} else {
				// Если результат < 0, смещаем указатель влево, увеличивая счетчик
				left++
			}
		}
	}

	return result
}
func main() {
	// fmt.Println(threeSum([]int{0, 0, 0, 0})) // [[0,0,0]]
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	// fmt.Println(threeSum([]int{0, 1, 1}))             // []
	// fmt.Println(threeSum([]int{0, 0, 0}))             // [[0,0,0]]
}
