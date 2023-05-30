package main

import (
	"fmt"
	"sort"
	"strings"
)

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

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string, len(strs))

	for _, key := range strs {
		resultKey := sortString(key)
		if _, ok := result[resultKey]; !ok {
			result[resultKey] = []string{}
		}

		result[resultKey] = append(result[resultKey], key)
	}

	res := make([][]string, 0)
	for _, k := range result {
		res = append(res, k)
	}

	return res
}

func sortString(s string) string {
	strSlice := strings.Split(s, "")
	sort.Strings(strSlice)
	return strings.Join(strSlice, "")
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // [["bat"],["nat","tan"],["ate","eat","tea"]]
}
