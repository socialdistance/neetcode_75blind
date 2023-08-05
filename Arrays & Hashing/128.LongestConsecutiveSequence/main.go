package main

import (
	"fmt"
)

// 128. Longest Consecutive Sequence

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

func longestConsecutive(nums []int) int {
	set := map[int]bool{}

	for _, num := range nums {
		set[num] = true
	}

	// Заводим счетчик для сохранения нашей последовательности
	res := 0

	for _, num := range nums {
		// Проверяем левого соседа, если соседа слева нет, значит мы предполагаем,
		// что это начало последовательности
		// 100 - 1 = 99
		if set[num-1] {
			continue
		}

		// Заводим две переменные
		// sequence -> для сохранения в цикле последовательности
		// temp -> переменная, в которой хранится значение num+1(100+1=101),
		// и если 101 есть в нашем слайсе, значит мы увеличиваем нашу
		// последовательность(sequence + 1), а так же temp+1
		sequence := 1
		temp := num + 1

		for set[temp] {
			sequence++
			temp++
		}

		// Если последовательность(sequence) > res, то присваиваем res новую максимальную
		// последовательность
		if sequence > res {
			res = sequence
		}
	}

	return res
}

func main() {
	fmt.Println(longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})) // 7
	fmt.Println(longestConsecutive([]int{0, -1}))                             // 2
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))              // 4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))      // 9
}
