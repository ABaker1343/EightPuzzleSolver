package main

import "fmt"

func main() {

	fmt.Println("input your starting node")
	var input [9]int
	var zeroIndex int
	for i := range input {
		fmt.Scanln(&input[i])
		if input[i] == 0 {
			zeroIndex = i
		}
	}
	fmt.Println("starting search...")

	initialNode := Node{
		value:     input,
		prev:      nil,
		zeroIndex: zeroIndex,
	}

	// initialNode := Node{
	// 	value:     [9]int{7, 2, 4, 5, 0, 6, 8, 3, 1},
	// 	prev:      nil,
	// 	zeroIndex: 4,
	// }

	goalNode := Node{
		value:     [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		zeroIndex: 8,
		prev:      nil,
	}

	goalNode = Node{
		value:     [9]int{1, 2, 3, 4, 5, 6, 7, 8, 0},
		zeroIndex: 7,
		prev:      nil,
	}

	//solveBoardItterator(initialNode, goalNode)
	solveBoard(initialNode, goalNode)
}
