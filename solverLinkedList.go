package main

import (
	"fmt"
	"time"
)

type NodeList struct {
	value *Node
	next  *NodeList
	last  *NodeList
}

func (nl *NodeList) append(node *Node) {
	newNodeList := &NodeList{
		value: node,
		next:  nil,
		last:  nil,
	}
	nl.last.next = newNodeList
	nl.last = newNodeList
}

type NodeListItterator struct {
	curr *NodeList
	list *NodeList
}

func (nli *NodeListItterator) Set(nodelist *NodeList) {
	nli.curr = nil
	nli.list = nodelist
}

func (nli *NodeListItterator) Next() bool {
	if nli.curr == nil {
		nli.curr = nli.list
		return true
	} else if nli.curr.next != nil {
		nli.curr = nli.curr.next
		return true
	}

	return false

}

func (nli *NodeListItterator) Get() *Node {
	return nli.curr.value
}

func solveBoardItterator(initialBoard Node, goalNode Node) {

	searchSpace := NodeList{
		value: &initialBoard,
		next:  nil,
	}
	searchSpace.last = &searchSpace

	found := false

	var finalNode *Node

	startTime := time.Now().Unix()

	for true {
		go func() {
			itt := NodeListItterator{}
			itt.Set(&searchSpace)
			for itt.Next() {
				if isSameBoard(&goalNode, itt.Get()) {
					found = true
					finalNode = itt.Get()
					break
				}
			}
		}()

		if !found {
			nextNode := getNextNodeItterator(&searchSpace, &goalNode)
			searchSpace.append(&nextNode)
		} else {
			break
		}
	}

	endTime := time.Now().Unix()

	finalNode = searchSpace.last.value

	printNodeTrace(finalNode)
	fmt.Printf("time elapsed %d, node depth: %d", endTime-startTime, getNodeDepth(finalNode))

}

func getNextNodeItterator(searchSpace *NodeList, goalNode *Node) Node {

	var bestEval int = 100000
	var bestNode Node

	itt := NodeListItterator{}
	itt.Set(searchSpace)

	for itt.Next() {
		n := itt.Get()
		for _, swapIndex := range getValidSwaps(n.zeroIndex) {
			newNode := swapWithZero(n, swapIndex)
			if !isInList(searchSpace, &newNode) {
				val := evaluate(&newNode, goalNode)
				if val < bestEval {
					bestNode = newNode
					bestEval = val
				}
			}
		}
	}

	return bestNode

}

func isInList(slice *NodeList, board *Node) bool {
	itt := NodeListItterator{}
	itt.Set(slice)
	for itt.Next() {
		if itt.Get().value == board.value {
			return true
		}
	}
	return false
}
