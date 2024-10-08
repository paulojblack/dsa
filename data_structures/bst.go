package ds

type BSTNode struct {
	left  *BSTNode
	right *BSTNode
	val   int
}

func preorder(root *BSTNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.val)
	preorder(root.left, result)
	preorder(root.right, result)

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

func find(root *BSTNode, val int) *BSTNode {
	if root == nil {
		return nil
	}

	if root.val == val {
		return root
	}

	if root.val > val {
		return find(root.left, val)
	}

	// Effectively <= case but should expect no == case
	return find(root.right, val)
}

// func delete(root *BSTNode, val int) bool {
// 	if root == nil {
// 		// Attempting to delete node that doesn't exist
// 		return false
// 	}

// 	if root.left != nil && root.left.val == val {

// 	}
// }
