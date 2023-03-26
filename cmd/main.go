package main

import (
	"github.com/afmahmuda/hungarian_algorithm/logic"
	"github.com/afmahmuda/hungarian_algorithm/types"
)

func main() {

	// worker := 4
	// job := 4
	// var costs []int
	// costs = []int{1, 2, 3, 4, 8, 7, 6, 5, 9, 12, 11, 10, 16, 14, 15, 13}
	// costs = []int{1, 2, 3, 4, 8, 7, 6, 5, 9, 12, 11, 10, 16, 14, 15, 16}
	// costs = []int{2, 2, 1, 3, 2, 1, 2, 1, 2, 1, 2, 3, 2, 1, 2, 3}
	// costs = []int{44, 68, 66, 21, 54, 70, 47, 97, 52, 9, 36, 60, 87, 89, 27, 19}

	worker := 10
	job := 10
	// costs := []int{
	// 	60, 4, 63, 92, 59, 36, 14, 50, 63, 65,
	// 	83, 87, 11, 61, 64, 6, 20, 40, 95, 15,
	// 	76, 71, 41, 83, 9, 53, 76, 26, 18, 82,
	// 	62, 59, 87, 45, 8, 93, 4, 83, 46, 57,
	// 	30, 8, 56, 2, 77, 17, 23, 58, 47, 46,
	// 	58, 33, 28, 83, 63, 41, 64, 95, 31, 63,
	// 	21, 43, 59, 99, 67, 39, 80, 48, 98, 86,
	// 	52, 56, 86, 94, 40, 28, 32, 18, 1, 84,
	// 	2, 33, 94, 94, 89, 78, 68, 12, 27, 12,
	// 	39, 77, 75, 63, 66, 27, 73, 19, 67, 9,
	// }
	costs := []int{
		68, 82, 68, 59, 47, 54, 35, 54, 17, 86,
		3, 77, 27, 6, 48, 35, 64, 26, 28, 2,
		54, 35, 85, 56, 54, 75, 23, 71, 77, 34,
		10, 44, 78, 3, 10, 62, 1, 64, 37, 68,
		25, 23, 33, 93, 80, 60, 85, 55, 84, 63,
		61, 97, 37, 6, 3, 88, 14, 73, 70, 29,
		33, 38, 56, 25, 25, 25, 8, 26, 40, 94,
		77, 17, 54, 16, 20, 8, 30, 54, 3, 22,
		55, 85, 29, 84, 74, 53, 16, 54, 61, 79,
		50, 34, 5, 44, 84, 95, 86, 2, 10, 91,
	}

	costMatrix, err := types.NewCostMatrix(worker, job, costs)
	if err != nil {
		panic(err.Error())
	}
	println("cost matrix:")
	println(costMatrix.String(), "\n")

	solution, _ := logic.Solve(*costMatrix)

	println("cost matrix:")
	println(costMatrix.String(), "\n")
	println("result matrix:")
	println(solution.String(), "\n")
	assignmentMap := logic.Translate(*costMatrix, solution)
	println("assignment mapping:")
	println(assignmentMap.String())
	println("with total cost: ", assignmentMap.TotalCost())
}
