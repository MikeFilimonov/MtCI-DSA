package graphs

import "fmt"

type Graph struct {
	// undirected and unweighted, using adjacency list
	numberOfNodes int
	adjacentList  map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		numberOfNodes: 0,
		adjacentList:  map[int][]int{},
	}
}

func (g *Graph) AddVertex(n int) {

	// if len(g.adjacentList) > 0 {
	// 	neighbouringPosition := len(g.adjacentList) - 1
	// 	g.adjacentList[neighbouringPosition] =
	// 		append(g.adjacentList[neighbouringPosition], n)
	// }
	g.adjacentList[n] = []int{}
	g.numberOfNodes++
}

func (g *Graph) AddEdge(nodeA, nodeB int) {
	//undirected graph
	g.adjacentList[nodeA] = append(g.adjacentList[nodeA], nodeB)
	g.adjacentList[nodeB] = append(g.adjacentList[nodeB], nodeA)
}

func (g *Graph) ShowConnections() {

	for i := 0; i < g.numberOfNodes; i++ {
		fmt.Println(fmt.Sprintf("%d --> %v", i, g.adjacentList[i]))
	}

}

func samples() {

	// 	  (2)--(0)
	// 	  /\
	//   /  \
	// (1)---(3)

	// edge list
	g := [][]int{
		{0, 2},
		{2, 3},
		{2, 1},
		{1, 3},
	}
	fmt.Println(g)

	// adjacent list
	g = [][]int{
		{2},       // zero item is connected to 2
		{2, 3},    // node one is connected to 2 and 3
		{0, 1, 3}, // node two is connected to 0, 1 and 3
		{1, 2},    // node three is connected with 1 and 2
	}

	// adjacent matrix
	g = [][]int{
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{1, 1, 0, 1},
		{0, 1, 1, 0},
	}

}
