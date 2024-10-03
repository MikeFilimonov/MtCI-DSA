package graphs

import "fmt"

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
