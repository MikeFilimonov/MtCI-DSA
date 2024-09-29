package trees

import (
	"fmt"
)

type BinaryTreeNode struct {
	data               int
	leftSideOfTheTree  *BinaryTreeNode
	rightSideOfTheTree *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func NewBinaryTreeNode(value int) *BinaryTreeNode {

	return &BinaryTreeNode{
		data:               value,
		leftSideOfTheTree:  nil,
		rightSideOfTheTree: nil,
	}

}

func NewBinaryTree(value int) *BinaryTree {

	return &BinaryTree{
		root: NewBinaryTreeNode(value),
	}
}

func (t *BinaryTree) Root() *BinaryTreeNode {
	return t.root
}

func (t *BinaryTree) Insert(value int) {

	newNode := NewBinaryTreeNode(value)

	if t.root == nil {
		t.root = newNode
		return
	}

	currentNode := t.root

	for {

		nextNode := currentNode.leftSideOfTheTree
		addToTheLeft := true

		if newNode.data > currentNode.data {
			nextNode = currentNode.rightSideOfTheTree
			addToTheLeft = false
		}

		if nextNode == nil {
			if addToTheLeft {
				currentNode.leftSideOfTheTree = newNode
			} else {
				currentNode.rightSideOfTheTree = newNode
			}
			return
		} else {
			currentNode = nextNode
		}

	}

}

func (t *BinaryTree) Lookup(value int) *BinaryTreeNode {

	currentNode := t.root
	if currentNode == nil {
		return nil
	}

	for currentNode != nil {

		if value == currentNode.data {
			return currentNode
		} else {

			if value > currentNode.data {
				currentNode = currentNode.rightSideOfTheTree
			} else {
				currentNode = currentNode.leftSideOfTheTree
			}

		}

	}
	return nil
}

func (t *BinaryTree) Traverse(node *BinaryTreeNode) *BinaryTreeNode {

	tree := node

	var leftBranchData int
	if node.leftSideOfTheTree != nil {
		leftBranchData = node.leftSideOfTheTree.data
	}

	var rightBranchData int

	if node.rightSideOfTheTree != nil {
		rightBranchData = node.rightSideOfTheTree.data
	}

	fmt.Printf("node: %d, left: %d, right: %d\n", node.data, leftBranchData, rightBranchData)

	if node.leftSideOfTheTree == nil {
		tree.leftSideOfTheTree = nil
	} else {
		tree.leftSideOfTheTree = t.Traverse(node.leftSideOfTheTree)
	}

	if node.rightSideOfTheTree == nil {
		tree.rightSideOfTheTree = nil
	} else {
		tree.rightSideOfTheTree = t.Traverse(node.rightSideOfTheTree)
	}

	return tree

}

func (t *BinaryTree) Remove(value int) {

	currentNode := t.root

	nodeToRemove := t.Lookup(value)
	if nodeToRemove == nil {
		return
	} else {

	}

}
