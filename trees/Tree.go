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

func (t *BinaryTree) RemoveLikeABoss(value int) {

	if t.root == nil {
		return
	}

	currentNode := t.root

	for currentNode != nil {

		previousNode := currentNode

		if currentNode.data > value {
			currentNode = currentNode.rightSideOfTheTree
		} else if currentNode.data < value {
			currentNode = currentNode.leftSideOfTheTree
		} else {

			// no right children
			if currentNode.rightSideOfTheTree == nil {
				if previousNode == nil {
					t.root = currentNode.leftSideOfTheTree
				} else {

					if currentNode.data < previousNode.data {
						previousNode.leftSideOfTheTree = currentNode.leftSideOfTheTree
					} else if currentNode.data > previousNode.data {
						previousNode.rightSideOfTheTree = currentNode.leftSideOfTheTree
					}
				}

				// right child without a left child
			} else if currentNode.rightSideOfTheTree.leftSideOfTheTree == nil {

				currentNode.rightSideOfTheTree.leftSideOfTheTree = currentNode.leftSideOfTheTree
				if previousNode == nil {
					t.root = currentNode.rightSideOfTheTree
				} else {

					if currentNode.data < previousNode.data {
						previousNode.leftSideOfTheTree = currentNode.rightSideOfTheTree
					} else if currentNode.data > previousNode.data {
						previousNode.rightSideOfTheTree = currentNode.rightSideOfTheTree
					}
				}

				// right child that has a left child
			} else {

				leftmost := currentNode.rightSideOfTheTree.leftSideOfTheTree
				leftmostParent := currentNode.rightSideOfTheTree
				for leftmost.leftSideOfTheTree != nil {
					leftmostParent = leftmost
					leftmost = leftmost.leftSideOfTheTree
				}

				// parent's left subtree is now the leftmost's right subtree
				leftmostParent.leftSideOfTheTree = leftmost.rightSideOfTheTree
				leftmost.leftSideOfTheTree = currentNode.leftSideOfTheTree
				leftmost.rightSideOfTheTree = currentNode.rightSideOfTheTree

				if previousNode == nil {
					t.root = leftmost
				} else {
					if currentNode.data < previousNode.data {
						previousNode.leftSideOfTheTree = leftmost
					} else if currentNode.data > previousNode.data {
						previousNode.rightSideOfTheTree = leftmost
					}
				}
				return

			}

		}

	}

}

func (t *BinaryTree) Remove(value int) {

	if t.root == nil {
		return
	}

	currentNode := t.root
	var previousNode *BinaryTreeNode
	previousNode = nil

	for currentNode != nil {

		if currentNode.data == value {
			break
		}

		previousNode = currentNode
		if currentNode.data > value {
			currentNode = currentNode.rightSideOfTheTree
		} else {
			currentNode = currentNode.leftSideOfTheTree
		}

	}

	var newNode *BinaryTreeNode
	newNode = nil
	if currentNode.rightSideOfTheTree != nil {

		newNode = currentNode.rightSideOfTheTree

		if newNode.leftSideOfTheTree != nil {
			newNode = newNode.leftSideOfTheTree
		}

	} else if newNode != nil {
		newNode = newNode.leftSideOfTheTree
	}

	if newNode != nil {

		if newNode.data > previousNode.data {
			previousNode.rightSideOfTheTree = newNode
		} else {
			previousNode.leftSideOfTheTree = newNode
		}

	}
}
