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

func getMetadataSum(node Node) int {
	metadataSum := 0

	for _,metadata := range node.metadata {
		metadataSum += metadata
	}

	for _,child := range node.children {
		metadataSum += getMetadataSum(child)
	}

	return metadataSum
}

func main() {
	inputs := au.ReadInputAsSingleLineNumberArray("08")
	// inputs := []int { 2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2 }

	node,_ := parseNode(inputs, 0)

	metadataSum := getMetadataSum(node)
	fmt.Println(metadataSum)
}
