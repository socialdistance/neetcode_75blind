package main

import "fmt"

// Given the root of a binary tree, invert the tree, and return its root.

// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) inOrder(root *TreeNode) {
	if root != nil {
		fmt.Println("data:", root.Val)
		root.inOrder(root.Left)
		root.inOrder(root.Right)
	}
}

func invertTree(root *TreeNode) *TreeNode {
	var result []*TreeNode

	if root == nil {
		return nil
	}

	result = append(result, root)

	for len(result) > 0 {
		rootNode := result[0]
		result = result[1:]

		if rootNode.Left != nil {
			result = append(result, rootNode.Left)
		}

		if rootNode.Right != nil {
			result = append(result, rootNode.Right)
		}

		newNode := rootNode.Left

		rootNode.Left = rootNode.Right
		rootNode.Right = newNode
	}

	return root
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(invertTree(root))
}
