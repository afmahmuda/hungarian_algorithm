package types

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type SquareMatrix[T ~int] struct {
	dimension int
	elements  []T
}

func (sm *SquareMatrix[T]) String() string {
	lines := []string{}

	maxNum := T(0)
	for i := 0; i < sm.dimension*sm.dimension; i++ {
		maxNum = T(int(math.Max(float64(maxNum), float64(sm.elements[i]))))
	}

	maxDigitElement := len(fmt.Sprint(maxNum))
	maxDigitDimension := len(fmt.Sprint(sm.dimension))

	headers := []string{}
	for i := 0; i < sm.dimension; i++ {
		headers = append(headers, fmt.Sprintf("%*d", maxDigitElement, i+1))
	}

	lines = append(lines, fmt.Sprintf("%*s   ", maxDigitDimension, "")+strings.Join(headers, "  "))

	for row := 0; row < sm.dimension; row++ {
		entry := []string{}
		for col := 0; col < sm.dimension; col++ {
			entry = append(entry, fmt.Sprintf("%*d", maxDigitElement, sm.Get(row, col)))
		}
		lines = append(lines, fmt.Sprintf("%*d [ %s ]", maxDigitDimension, row+1, strings.Join(entry, ", ")))
	}

	return strings.Join(lines, "\n")
}

func NewEmptySquareMatrix[T ~int](dimension int) *SquareMatrix[T] {
	return &SquareMatrix[T]{
		dimension: dimension,
		elements:  make([]T, dimension*dimension),
	}

}

// Set... row and col is zero based
func (sm *SquareMatrix[T]) Set(row, col int, value T) {
	idx := row*sm.dimension + col
	sm.elements[idx] = value
}

// Get... row and col is zero based
func (sm *SquareMatrix[T]) Get(row, col int) T {
	idx := row*sm.dimension + col
	return sm.elements[idx]
}

func NewSquareMatrix[T ~int](rowNum, colNum int, elements []T) (*SquareMatrix[T], error) {
	if len(elements) != rowNum*colNum {
		return nil, errors.New("len(costs) != workerCount*jobCount")
	}

	dimension := int(math.Max(float64(rowNum), float64(colNum)))

	result := NewEmptySquareMatrix[T](dimension)

	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			// index := row*dimension + col
			// result.elements[index] = elements[row*colNum+col]
			result.Set(row, col, elements[row*colNum+col])
		}
	}

	return result, nil
}

func NewCostMatrix[T ~int](workerCount, jobCount int, costs []T) (*SquareMatrix[T], error) {

	return NewSquareMatrix(workerCount, jobCount, costs)

}
