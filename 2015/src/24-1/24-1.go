package main

import (
	"au"
	"fmt"
)

func getTargetSum(inputs []int) int {
	sum := 0

	for _,input := range inputs {
		sum += input
	}

	return sum / 3
}

func hasGroup(inputs []int, groupSize int, groupSum int) bool {
	if (groupSize == 2) {
		for i := 0; i < len(inputs); i++ {
			for j := i + 1; j < len(inputs); j++ {
				if inputs[i] + inputs[j] == groupSum {
					return true
				}
			}
		}
	} else {
		for i := 0; i < len(inputs); i++ {
			subGroupSum := groupSum - inputs[i] 

			if subGroupSum > 0 {
				subInputs := append(inputs[:i], inputs[i+1:]...)
				subGroupSize := groupSize - 1 

				if hasGroup(subInputs, subGroupSize, subGroupSum) {
					return true
				}
			}
		}
	}

	return false
}

func getGroups(inputs []int, groupSize int, groupSum int) []map[int] bool {
	groups := make([]map[int] bool, 0)

//	fmt.Println("Inputs: ", len(inputs), groupSize, groupSum)

	if (groupSize == 2) {
		for i := 0; i < len(inputs); i++ {
			for j := i + 1; j < len(inputs); j++ {
				if inputs[i] + inputs[j] == groupSum {
					group := make(map[int] bool)
					group[inputs[i]] = true
					group[inputs[j]] = true
					groups = append(groups, group)
				}
			}
		}
	} else if groupSize > 0 {
		for i := 0; i < len(inputs); i++ {
			subGroupSum := groupSum - inputs[i] 

			if subGroupSum > 0 {
				subInputs := make([]int, 0)
				for j := 0; j < len(inputs); j++ {
					if i != j {
						subInputs = append(subInputs, inputs[j])
					}
				}
				subGroupSize := groupSize - 1 

				for _, subGroup := range getGroups(subInputs, subGroupSize, subGroupSum) {
					subGroup[inputs[i]] = true
					groups = append(groups, subGroup)
				}
			}
		}
	}

	return groups
}

func getSubInput(inputs []int, excludes map [int] bool) []int {
	subInput := make([]int, 0)
	
	for _,input := range inputs {
		_,ok := excludes[input]
		if !ok {
			subInput = append(subInput, input)
		}
	}

	return subInput
}

func getQuantumEntanglement(group map [int] bool) int {
	qe := 1

	for product := range group {
		qe *= product
	}

	return qe
}

func hasValidGroupTwoAndThree(inputs2 []int, targetSum int) bool {
	for j := 1; j < len(inputs2) - 1; j++ {
		group2Size := j

		for _,group2 := range getGroups(inputs2, group2Size, targetSum) {
			inputs3 := getSubInput(inputs2, group2)

			group3Size := len(inputs2) - group2Size

			if hasGroup(inputs3, group3Size, targetSum) {
				return true
			}
		}
	}

	return false
}

func getBestQuantumEntanglement(inputs []int, targetSum int) int {
	bestQE := -1

	for i := 1; i < len(inputs); i++ {
		group1Size := i

		for _,group1 := range getGroups(inputs, group1Size, targetSum) {
			qe := getQuantumEntanglement(group1)
			if bestQE != -1 && qe >= bestQE {
				continue
			}

			inputs2 := getSubInput(inputs, group1)
			if hasValidGroupTwoAndThree(inputs2, targetSum) {
				bestQE = qe
			}
		}

		if bestQE != -1 {
			break;
		}
	}

	return bestQE
}

func main() {
	inputs := au.ReadInputAsNumberArray("24")
	// inputs = []int { 1, 2, 3, 4, 5, 7, 8, 9, 10, 11 }

	targetSum := getTargetSum(inputs)

	candidates := getBestQuantumEntanglement(inputs, targetSum)

	fmt.Println(candidates)
}
