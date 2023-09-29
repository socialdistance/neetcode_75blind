package main

import "fmt"

// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

// A leaf is a node with no children.

// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// Output: true
// Explanation: The root-to-leaf path with the target sum is shown.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	curSum := 0

	return dfs(root, curSum)
}

func dfs(root *TreeNode, curSum int) bool {
	if root == nil {
		return false
	}

	curSum += root.Val

	if root.Left == nil && root.Right == nil {
		fmt.Println(curSum)
	}

	return dfs(root.Left, curSum) || dfs(root.Right, curSum)
}

// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if hasChild(root) && targetSum == root.Val {
// 		return true
// 	}

// 	if hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) {
// 		return true
// 	}

// 	return false
// }

// func hasChild(root *TreeNode) bool {
// 	return root.Left == nil && root.Right == nil
// }

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	fmt.Println(hasPathSum(root, 22))
}
