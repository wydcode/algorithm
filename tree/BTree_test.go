package tree

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func initBTree() *BTree{
	node1 := &BTree{Val: 3}
	node2 := &BTree{Val: 4}
	node3 := &BTree{Val: 6}
	node4 := &BTree{Val: 8}
	node5 := &BTree{Val: 10}
	node6 := &BTree{Val: 11}

	btree := &BTree{Val: 2, Left: node1, Right: node2}
	btree.Left.Left = node3
	btree.Left.Right = node4

	btree.Right.Left = node5
	btree.Right.Right = node6

	return btree
}

func TestBTree_New(t *testing.T){
	//btree := initBTree()
	//btree.Preorder(btree)
	//btree.LevelOrder(btree)

	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	//postorder := []int{9,15,7,20,3}

	btree := &BTree{}
	tree := btree.BuildTree(preorder, inorder)
	//tree := btree.BuildTree(inorder, postorder)
	spew.Dump(tree)
}

func TestBTree_BuildTree(t *testing.T) {
	preorder := []int{3,9,20,15,7}
	tree := BuildBinaryTree(0, preorder)
	spew.Dump(tree)
}


