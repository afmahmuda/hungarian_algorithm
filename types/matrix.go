package types

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type SquareMatrix struct {
	dimension int
	elements  []int
}

// Set... row and col is zero based
func (sm *SquareMatrix) Set(row, col int, value int) {
	idx := row*sm.dimension + col
	sm.elements[idx] = value
}

func (sm *SquareMatrix) Dim() int {
	return sm.dimension
}

// Get... row and col is zero based
func (sm *SquareMatrix) Get(row, col int) int {
	idx := row*sm.dimension + col
	return sm.elements[idx]
}

func (sm *SquareMatrix) Clone() *SquareMatrix {
	elem := make([]int, sm.dimension*sm.dimension)

	copy(elem, sm.elements)
	return &SquareMatrix{
		dimension: sm.dimension,
		elements:  elem,
	}
}

func (sm *SquareMatrix) String() string {
	lines := []string{}

	maxNum := int(0)
	for i := 0; i < sm.dimension*sm.dimension; i++ {
		maxNum = int(int(math.Max(float64(maxNum), float64(sm.elements[i]))))
	}

	maxDigitElement := len(fmt.Sprint(maxNum))
	maxDigitDimension := len(fmt.Sprint(sm.dimension))

	headers := []string{}
	for i := 0; i < sm.dimension; i++ {
		headers = append(headers, fmt.Sprintf("%*d", maxDigitElement, i))
	}

	lines = append(lines, fmt.Sprintf("%*s   ", maxDigitDimension, "")+strings.Join(headers, "  "))

	for row := 0; row < sm.dimension; row++ {
		entry := []string{}
		for col := 0; col < sm.dimension; col++ {
			entry = append(entry, fmt.Sprintf("%*d", maxDigitElement, sm.Get(row, col)))
		}
		lines = append(lines, fmt.Sprintf("%*d [ %s ]", maxDigitDimension, row, strings.Join(entry, ", ")))
	}

	return strings.Join(lines, "\n")
}

func NewSquareMatrix(dimension int) *SquareMatrix {
	return &SquareMatrix{
		dimension: dimension,
		elements:  make([]int, dimension*dimension),
	}

}

func NewCostMatrix(workerCount, jobCount int, costs []int) (*SquareMatrix, error) {

	rowNum, colNum, elements := workerCount, jobCount, costs

	if len(elements) != rowNum*colNum {
		return nil, errors.New("len(costs) != workerCount*jobCount")
	}

	dimension := int(math.Max(float64(rowNum), float64(colNum)))

	result := NewSquareMatrix(dimension)

	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			result.Set(row, col, elements[row*colNum+col])
		}
	}

	return result, nil
}
