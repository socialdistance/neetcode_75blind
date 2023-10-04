package main

import (
	"fmt"
	"math"
)

// Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	minDiff := math.MaxUint32
	prev := -1

	dfs(root, &prev, &minDiff)

	return minDiff
}

func dfs(r *TreeNode, prev *int, minDiff *int) int {
	if r == nil {
		return 0
	}

	dfs(r.Left, prev, minDiff)

	if *prev != -1 && r.Val-*prev < *minDiff {
		*minDiff = r.Val - *prev
	}

	prev = &r.Val

	dfs(r.Right, prev, minDiff)

	return *minDiff
}

func main() {
	// root := &TreeNode{
	// 	Val: 4,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val:   1,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   3,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   6,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// }

	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   49,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(minDiffInBST(root1))
}
