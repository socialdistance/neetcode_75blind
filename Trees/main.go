package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) find(key int) *TreeNode {
	current := t

	for current.Val != key {
		if key < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}

		if current == nil {
			return nil
		}
	}

	return current
}

func (t *TreeNode) insert(id int) {
	current := t

	if id <= current.Val {
		if current.Left == nil {
			current.Left = &TreeNode{Val: id}
		} else {
			current.Left.insert(id)
		}
	} else {
		if current.Right == nil {
			current.Right = &TreeNode{Val: id}
		} else {
			current.Right.insert(id)
		}
	}
}

func (t *TreeNode) delete(id int) {

}

func (t *TreeNode) inOrder(root *TreeNode) {
	if root != nil {
		root.inOrder(root.Left)
		fmt.Println("data:", root.Val)
		root.inOrder(root.Right)
	}
}

func minimum(root *TreeNode) *TreeNode {
	var last *TreeNode
	for root != nil {
		last = root
		root = root.Left
	}

	return last
}

func maximum(root *TreeNode) *TreeNode {
	var last *TreeNode
	for root != nil {
		last = root
		root = root.Right
	}

	return last
}

func main() {
	root := new(TreeNode)

	root.insert(50)
	root.insert(25)
	root.insert(75)

	// root.inOrder(root)

	fmt.Println(maximum(root))

	// found := root.find(25)
	// fmt.Println(found)
}
