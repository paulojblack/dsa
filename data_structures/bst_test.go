package ds

import (
	"fmt"
	"testing"
)

func setupBST() *BSTNode {

	sample := []int{15, 2, 1, 5, 19, 2, 3, 4, 192, 4}

	root := &BSTNode{}
	for _, v := range sample {
		root = insert(root, v)
	}
	fmt.Println(root)
	return root
}

func TestBSTBasic(t *testing.T) {
	root := setupBST()

	walk(root)
	if 1 == 2 {
		t.Errorf("erorr you are a loser")
	}
}
