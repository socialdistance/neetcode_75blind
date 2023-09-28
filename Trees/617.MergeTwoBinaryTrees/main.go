package main

import "fmt"

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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root := &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  nil,
		Right: nil,
	}

	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)

	return root
}

func main() {
	// root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	mergeTrees(root1, root2)

}
