package main

// Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.
//
// Example 1:
//
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
// Example 2:
//
// Input: nums = [1,1]
// Output: [2]
func findDisappearedNumbers(nums []int) []int {
	N := len(nums)
	cnt := make([]int, N+1)
	for _, num := range nums {
		cnt[num] = 1
	}

	ans := []int{}

	for i := 1; i <= N; i++ {
		if cnt[i] == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}) // [5, 6]
	findDisappearedNumbers([]int{1, 1})                   // [2]
}
