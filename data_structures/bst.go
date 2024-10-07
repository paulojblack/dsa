package ds

import "fmt"

type BSTNode struct {
	left  *BSTNode
	right *BSTNode
	val   int
}

func walk(root *BSTNode) {
	if root == nil {
		return
	}

	fmt.Println(root.val)
	walk(root.left)
	walk(root.right)

}

func insert(root *BSTNode, val int) *BSTNode {
	if root == nil {
		root = &BSTNode{nil, nil, val}
	} else if root.val > val {
		root.left = insert(root.left, val)
	} else if root.val < val {
		root.right = insert(root.right, val)
	}

	return root
}
