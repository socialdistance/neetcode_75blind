package main

import "fmt"

//Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
//
//
//Example 1:
//
//Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
//Example 2:
//
//Input: strs = [""]
//Output: [[""]]
//
//Example 3:
//
//Input: strs = ["a"]
//Output: [["a"]]

// Разбить массив по две строки, передавать их и сделать так же, как с двумя строками, а после собрать это в слайс слайсов ?

func groupAnagrams(strs []string) [][]string {
	result := make(map[byte]int)

	for idx := 0; idx < len(strs); idx++ {
		//fmt.Println(strs[idx])
		//result[strs[idx]]++
		for i := 0; i < len(strs[idx]); i++ {
			fmt.Println(strs[idx][i]) // bytes
			result[strs[idx][i]]++
		}
	}

	fmt.Println(result)

	return [][]string{}
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // [["bat"],["nat","tan"],["ate","eat","tea"]]
}
