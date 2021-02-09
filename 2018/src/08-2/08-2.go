package main

import (
	"au"
	"fmt"
)

type Node struct {
	children []Node
	metadata []int
}

func parseNode(inputs []int, startIndex int) (Node, int) {
	index := startIndex

	childCount := inputs[index]
	index++

	metadataCount := inputs[index]
	index++

	children := []Node{}
	for i := 0; i < childCount; i++ {
		child,length := parseNode(inputs, index)

		children = append(children, child)
		index += length
	}

	metadata := inputs[index:index + metadataCount]
	index += metadataCount

	return Node {children, metadata}, index - startIndex
}

func getValue(node Node) int {
	if len(node.children) == 0 {
		metadataSum := 0

		for _,metadata := range node.metadata {
			metadataSum += metadata
		}

		return metadataSum
	} else {

		value := 0

		for _,metadata := range node.metadata {
			if metadata == 0 {
				continue
			}

			childIndex := metadata - 1
			if childIndex < len(node.children) {
				value += getValue(node.children[childIndex])
			}
		}

		return value
	}

	return 0
}

func main() {
	inputs := au.ReadInputAsSingleLineNumberArray("08")
	// inputs := []int { 2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2 }

	node,_ := parseNode(inputs, 0)

	value := getValue(node)
	fmt.Println(value)
}
