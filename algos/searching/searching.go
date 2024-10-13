package searching

import "MtCI-DSA/trees"

func BreadthFirstSearch() []int {

	bst := dummyTree()
	return bst.BreadthFirstSearch()

}

func RecursiveBreadthFirstSearch() []int {

	bst := dummyTree()
	queue := []*trees.BinaryTreeNode{}
	queue = append(queue, bst.Root())
	list := []int{}

	return bst.RecursiveBreadthFirstSearch(queue, list)

}

func dummyTree() *trees.BinaryTree {

	bst := trees.NewBinaryTree(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	return bst

}

func theory() {

	/* BFS vs DFS

	   //If you know a solution is not far from the root of the tree:
	   BFS


	   //If the tree is very deep and solutions are rare:
	   BFS as DFS will take took long


	   //If the tree is very wide:
		DFS as BFS will take too much space


	   //If solutions are frequent but located deep in the tree:
	   DFS


	   //Determining whether a path exists between two nodes:
	   DFS


	   //Finding the shortest path:
	   BFS

	*/

}
