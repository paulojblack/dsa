package ds

import (
	"testing"
)

func setupBST(sample []int) *BSTNode {
	var root *BSTNode
	for _, v := range sample {
		root = insert(root, v)
	}
	return root
}

func TestBSTBasic(t *testing.T) {
	sample := []int{15, 1, 5, 19, 2, 3, 4, 192}
	root := setupBST(sample)

	expected := []int{15, 1, 5, 2, 3, 4, 19, 192}
	var result []int
	preorder(root, &result)

	for k, _ := range result {
		if result[k] != expected[k] {
			t.Errorf("element %v of result does not equal that of expected. Result: %v, Expected: %v", k, result[k], expected[k])
		}
	}
}

func TestBSTFindWhenExists(t *testing.T) {
	sample := []int{15, 1, 5, 19, 2, 3, 4, 192}
	root := setupBST(sample)

	found := find(root, 19)

	if found == nil {
		t.Errorf("failed to find existing node")
	} else if found.val != 19 {
		t.Errorf("found the dang node but the dang value aint the right dang one")
	}
}

func TestBSTFindWhenNotExists(t *testing.T) {
	sample := []int{15, 1, 5, 19, 2, 3, 4, 192}
	root := setupBST(sample)

	found := find(root, 20)

	if found != nil {
		t.Errorf("found a node that wasnt there. spooky")
	}
}
