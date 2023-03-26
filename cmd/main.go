package main

import "github.com/afmahmuda/hungarian_algorithm/types"

func main() {

	worker := 5
	job := 4
	costs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	costMatrix, err := types.NewCostMatrix(worker, job, costs)
	if err != nil {
		panic(err.Error())
	}
	println("cost matrix:")
	println(costMatrix.String(), "\n")

	result := costMatrix.Clone()

	result.Set(0, 0, 99)

	println("result matrix:")
	println(result.String(), "\n")

}
