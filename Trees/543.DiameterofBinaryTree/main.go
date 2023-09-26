package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxLen := 0
	dfs(root, &maxLen)

	return maxLen
}

func dfs(r *TreeNode, maxLen *int) int {
	if r == nil {
		return 0
	}

	left := dfs(r.Left, maxLen)
	right := dfs(r.Right, maxLen)

	*maxLen = max(*maxLen, left+right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(diameterOfBinaryTree(root))
}
