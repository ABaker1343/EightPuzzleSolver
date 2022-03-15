package main

import (
	"fmt"
	"time"
)

type Node struct {
	value     [9]int
	prev      *Node
	zeroIndex int
}

func solveBoard(initialBoard Node, goalNode Node) {

	var searchSpace []Node = []Node{}

	searchSpace = append(searchSpace, initialBoard)

	var found bool = false

	var finalNode *Node

	startTime := time.Now().Unix()

	for true {
		go func() {
			for i := range searchSpace {
				if isSameBoard(&goalNode, &searchSpace[i]) {
					found = true
					finalNode = &searchSpace[i]
					break
				}
			}
		}()

		if !found {
			nextNode := getNextNode(searchSpace, &goalNode)
			searchSpace = append(searchSpace, nextNode)
		} else {
			break
		}
	}

	endTime := time.Now().Unix()

	printNodeTrace(finalNode)
	fmt.Printf("time elapsed %d, node depth: %d", endTime-startTime, getNodeDepth(finalNode))

}

func isSameBoard(b1 *Node, b2 *Node) bool {
	return b1.value == b2.value
}

func isInSlice(slice []Node, board *Node) bool {
	for i := range slice {
		if slice[i].value == board.value {
			return true
		}
	}
	return false
}

func getNextNode(searchSpace []Node, goalNode *Node) Node {

	var bestEval int = 100000
	var bestNode Node

	for i := range searchSpace {
		n := &searchSpace[i]
		for _, swapIndex := range getValidSwaps(n.zeroIndex) {
			newNode := swapWithZero(n, swapIndex)
			val := evaluate(&newNode, goalNode)
			if val < bestEval {
				if !isInSlice(searchSpace, &newNode) {
					bestNode = newNode
					bestEval = val
				}
			}
		}
	}

	return bestNode

}

func printNodeTrace(node *Node) {
	var currNode *Node = node

	for true {
		printNode(currNode)
		if currNode.prev != nil {
			currNode = currNode.prev
		} else {
			break
		}
	}
}

func printNode(node *Node) {
	for i := 0; i < 3; i++ {
		fmt.Print(node.value[i])
	}
	fmt.Print("\n")
	for i := 0; i < 3; i++ {
		fmt.Print(node.value[i+3])
	}
	fmt.Print("\n")
	for i := 0; i < 3; i++ {
		fmt.Print(node.value[i+6])
	}
	fmt.Print("\n\n")
}

func getNodeDepth(node *Node) int {
	var currNode *Node = node
	var depth int = -1

	for true {
		depth++
		if currNode.prev != nil {
			currNode = currNode.prev
		} else {
			break
		}
	}

	return depth
}

func evaluate(node *Node, goalNode *Node) int {
	var totalDistance int = 0

	for i := 0; i < len(goalNode.value); i++ {
		if node.value[i] != 0 {
			totalDistance += getDistanceFromFinalLocation(node.value[i], i, goalNode)
		}
	}

	return totalDistance + getNodeDepth(node)
}

func getDistanceFromFinalLocation(value int, index int, goalNode *Node) int {
	var targetIndex int = -1
	var distance int = 0

	for i := 0; i < len(goalNode.value); i++ {
		if value == goalNode.value[i] {
			targetIndex = i
		}
	}

	// you can move up/down and left/right but
	// you have to make sure that is represented
	// by the moves made in this function

	for index != targetIndex {
		//its its above / below
		if index <= targetIndex-3 {
			index += 3
		} else if index >= targetIndex+3 {
			index -= 3
		} else if index%3 == 1 {
			//you can move it both left and right
			if index <= targetIndex-2 {
				//move down
				index += 3
			} else if index >= targetIndex+2 {
				index -= 3
			}
			if index < targetIndex {
				index++
			} else {
				index--
			}
		} else if index%3 == 0 {
			//left side
			//you can just increment beacuse its not
			//above or below and its not in the right spot
			index++
		} else if index%3 == 2 {
			//right side
			//you can just decrement because its not
			//above or below and its not in the right spot
			index--
		}

		distance++
	}

	return distance

}

func getValidSwaps(zeroIndex int) []int {

	var validSwaps []int

	switch zeroIndex {
	case 0:
		validSwaps = []int{1, 3}
	case 1:
		validSwaps = []int{0, 2, 4}
	case 2:
		validSwaps = []int{1, 5}
	case 3:
		validSwaps = []int{0, 4, 6}
	case 4:
		validSwaps = []int{1, 3, 5, 7}
	case 5:
		validSwaps = []int{2, 4, 8}
	case 6:
		validSwaps = []int{3, 7}
	case 7:
		validSwaps = []int{4, 6, 8}
	case 8:
		validSwaps = []int{5, 7}
	default:
		panic("no Moves")
	}

	return validSwaps
}

func swapWithZero(baseNode *Node, index int) Node {
	tempNode := Node{
		value:     baseNode.value,
		zeroIndex: baseNode.zeroIndex,
		prev:      baseNode,
	}

	tempNode.value[tempNode.zeroIndex] = tempNode.value[index]
	tempNode.value[index] = 0
	tempNode.zeroIndex = index

	return tempNode
}
