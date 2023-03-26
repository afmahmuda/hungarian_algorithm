package logic

import (
	"math"
	"strconv"
	"strings"

	"github.com/afmahmuda/hungarian_algorithm/types"
)

type step int

const (
	step_1 step = iota + 1
	step_2
	step_3
	step_4
	step_5
	step_6
	step_done
)

type rowCoverage map[int]bool
type colCoverage map[int]bool

func (rc rowCoverage) String() string {
	if rc == nil {
		return "no row coverage"
	}

	items := []string{}
	for k := range rc {
		items = append(items, strconv.Itoa(k))
	}
	return "row covered: " + strings.Join(items, ", ")
}

func (cc colCoverage) String() string {
	if cc == nil {
		return "no col coverage"
	}
	items := []string{}
	for k := range cc {
		items = append(items, strconv.Itoa(k))
	}
	return "col covered: " + strings.Join(items, ", ")
}

func Solve(costMatrix types.SquareMatrix) (types.SquareMatrix, error) {

	step := step_1
	newCost := *costMatrix.Clone()
	maskMatrix := *types.NewSquareMatrix(newCost.Dim())

	var rowCov rowCoverage
	var colCov colCoverage

	var initial_row_smallest_value, initial_col_smallest_value int
loop:
	for {
		// time.Sleep(1 * time.Second)
		println("\nstep", step)
		switch step {
		case step_1:
			newCost, step = step1(newCost)
			println("minima cost matrix")
			println(newCost.String())
		case step_2:
			rowCov = rowCoverage{}
			maskMatrix, colCov, step = step2(newCost)
			println("init mask matrix")
			println(maskMatrix.String())
			println(rowCov.String())
			println(colCov.String())

		case step_3:
			colCov, step = step3(maskMatrix)
			println(rowCov.String())
			println("is solution found?", step == step_done)
		case step_4:
			maskMatrix, rowCov, colCov, initial_row_smallest_value, initial_col_smallest_value, step = step4(newCost, maskMatrix, rowCov, colCov)
			println("mask matrix")
			println(maskMatrix.String())
			println(rowCov.String())
			println(colCov.String())
		case step_5:
			maskMatrix, rowCov, colCov, step = step5(newCost, maskMatrix, rowCov, colCov, initial_row_smallest_value, initial_col_smallest_value)
			println("mask matrix")
			println(maskMatrix.String())
			println(rowCov.String())
			println(colCov.String())
		case step_6:
			newCost, step = step6(newCost, maskMatrix, rowCov, colCov)
			println("altered cost matrix")
			println(maskMatrix.String())
		default:
			println("done")
			break loop
		}
	}

	return maskMatrix, nil
}

func Translate(costMatrix, solutionMatrix types.SquareMatrix) types.AssignmentMap {
	mapping := types.AssignmentMap{}
	for row := 0; row < solutionMatrix.Dim(); row++ {
		for col := 0; col < solutionMatrix.Dim(); col++ {
			if solutionMatrix.Get(row, col) == 1 {
				mapping[types.Assignee(row)] = types.Assignment{
					Assignee: types.Assignee(row),
					Job:      types.Job(col),
					Cost:     types.Cost(costMatrix.Get(row, col)),
				}
			}
		}
	}
	return mapping
}

func step1(costMatrix types.SquareMatrix) (types.SquareMatrix, step) {
	result := costMatrix.Clone()

	rowCount, colCount := result.Dim(), result.Dim()
	for row := 0; row < rowCount; row++ {
		minVal := math.MaxInt
		for col := 0; col < colCount; col++ {
			minVal = int(math.Min(float64(minVal), float64(result.Get(row, col))))
		}
		for col := 0; col < colCount; col++ {
			result.Set(row, col, result.Get(row, col)-minVal)
		}
	}

	for col := 0; col < rowCount; col++ {
		minVal := math.MaxInt
		for row := 0; row < colCount; row++ {
			minVal = int(math.Min(float64(minVal), float64(result.Get(row, col))))
		}
		for row := 0; row < colCount; row++ {
			result.Set(row, col, result.Get(row, col)-minVal)
		}
	}

	return *result, step_2
}

func step2(costMatrix types.SquareMatrix) (types.SquareMatrix, colCoverage, step) {
	result := types.NewSquareMatrix(costMatrix.Dim())

	rowCoveredMap := rowCoverage{}
	colCoveredMap := colCoverage{}
	for col := 0; col < costMatrix.Dim(); col++ {
		for row := 0; row < costMatrix.Dim(); row++ {
			if costMatrix.Get(row, col) != 0 {
				continue
			}
			if rowCoveredMap[row] {
				continue
			}
			if colCoveredMap[col] {
				continue
			}
			rowCoveredMap[row] = true
			colCoveredMap[col] = true
			result.Set(row, col, 1)
		}
	}

	return *result, colCoveredMap, step_3
}

func step3(maskMatrix types.SquareMatrix) (colCoverage, step) {

	colCoveredMap := colCoverage{}

	for col := 0; col < maskMatrix.Dim(); col++ {
		for row := 0; row < maskMatrix.Dim(); row++ {
			if maskMatrix.Get(row, col) == 1 {
				colCoveredMap[col] = true
				break
			}
		}
	}

	if len(colCoveredMap) < maskMatrix.Dim() {
		return colCoveredMap, step_4
	}

	return colCoveredMap, step_done
}

type rowColPair struct {
	row, col int
}

func step4(costMatrix, maskMatrix types.SquareMatrix, rowCoveredMap rowCoverage, colCoveredMap colCoverage) (types.SquareMatrix, rowCoverage, colCoverage, int, int, step) {

	if len(colCoveredMap) == costMatrix.Dim() {
		return maskMatrix, nil, nil, 0, 0, step_done
	}

	for {
		found, row, col := getUncoveredZero(costMatrix, maskMatrix, rowCoveredMap, colCoveredMap)
		if !found {
			return maskMatrix, rowCoveredMap, colCoveredMap, 0, 0, step_6
		}

		maskMatrix.Set(row, col, 2)
		if isStarInRow(maskMatrix, row) {
			col = starInRow(maskMatrix, row)
			rowCoveredMap[row] = true
			delete(colCoveredMap, col)
			continue
		}

		return maskMatrix, rowCoveredMap, colCoveredMap, row, col, step_5

	}

}

func step5(costMatrix, maskMatrix types.SquareMatrix, rowCoveredMap rowCoverage, colCoveredMap colCoverage, initial_row_smallest_value, initial_col_smallest_value int) (types.SquareMatrix, rowCoverage, colCoverage, step) {
	paths := []rowColPair{}

	paths = append(paths, rowColPair{initial_row_smallest_value, initial_col_smallest_value})

	for {
		col := paths[len(paths)-1].col

		found := isStarInCol(maskMatrix, col)
		if found {
			r := starInCol(maskMatrix, col)
			paths = append(paths, rowColPair{r, col})

			c := primeInRow(maskMatrix, r)
			paths = append(paths, rowColPair{r, c})
			continue
		}
		break
	}

	maskMatrix = augmentPath(maskMatrix, paths)
	maskMatrix = erasePrimes(maskMatrix)

	return maskMatrix, rowCoverage{}, colCoverage{}, step_3
}

func step6(costMatrix, maskMatrix types.SquareMatrix, rowCoveredMap rowCoverage, colCoveredMap colCoverage) (types.SquareMatrix, step) {
	newCost := costMatrix.Clone()

	minVal := findSmallest(costMatrix, rowCoveredMap, colCoveredMap)
	for row := 0; row < maskMatrix.Dim(); row++ {
		for col := 0; col < maskMatrix.Dim(); col++ {
			newValue := newCost.Get(row, col)
			if rowCoveredMap[row] {
				newValue += minVal
			}
			if !colCoveredMap[col] {
				newValue -= minVal
			}
			newCost.Set(row, col, newValue)
		}
	}

	return *newCost, step_4
}

func getUncoveredZero(costMatrix, maskMatrix types.SquareMatrix, rowCoveredMap, colCoveredMap map[int]bool) (bool, int, int) {
	for row := 0; row < maskMatrix.Dim(); row++ {
		for col := 0; col < maskMatrix.Dim(); col++ {

			if costMatrix.Get(row, col) != 0 {
				continue
			}
			if maskMatrix.Get(row, col) == 1 {
				continue
			}
			if rowCoveredMap[row] {
				continue
			}
			if colCoveredMap[col] {
				continue
			}
			return true, row, col
		}
	}
	return false, -1, -1
}

func isStarInRow(maskMatrix types.SquareMatrix, row int) bool {
	for col := 0; col < maskMatrix.Dim(); col++ {
		if maskMatrix.Get(row, col) == 1 {
			return true
		}
	}
	return false
}

func starInRow(maskMatrix types.SquareMatrix, row int) int {
	for col := 0; col < maskMatrix.Dim(); col++ {
		if maskMatrix.Get(row, col) == 1 {
			return col
		}
	}
	return -1
}

func primeInRow(maskMatrix types.SquareMatrix, row int) int {
	for col := 0; col < maskMatrix.Dim(); col++ {
		if maskMatrix.Get(row, col) == 2 {
			return col
		}
	}
	return -1
}

func isStarInCol(maskMatrix types.SquareMatrix, col int) bool {
	for row := 0; row < maskMatrix.Dim(); row++ {
		if maskMatrix.Get(row, col) == 1 {
			return true
		}
	}
	return false
}

func starInCol(maskMatrix types.SquareMatrix, col int) int {
	for row := 0; row < maskMatrix.Dim(); row++ {
		if maskMatrix.Get(row, col) == 1 {
			return row
		}
	}
	return -1
}

func augmentPath(maskMatrix types.SquareMatrix, paths []rowColPair) types.SquareMatrix {
	newmaskMatrix := maskMatrix.Clone()

	for _, v := range paths {
		val := 1
		if maskMatrix.Get(v.row, v.col) == 1 {
			val = 0
		}
		newmaskMatrix.Set(v.row, v.col, val)

	}

	return *newmaskMatrix

}

func erasePrimes(maskMatrix types.SquareMatrix) types.SquareMatrix {
	newmaskMatrix := maskMatrix.Clone()
	for row := 0; row < maskMatrix.Dim(); row++ {
		for col := 0; col < maskMatrix.Dim(); col++ {
			if newmaskMatrix.Get(row, col) == 2 {
				newmaskMatrix.Set(row, col, 0)
			}
		}
	}
	return *newmaskMatrix
}

func findSmallest(costMatrix types.SquareMatrix, rowCov rowCoverage, colCov colCoverage) int {
	minVal := math.MaxInt
	for row := 0; row < costMatrix.Dim(); row++ {
		for col := 0; col < costMatrix.Dim(); col++ {
			if rowCov[row] {
				continue
			}
			if colCov[col] {
				continue
			}
			if costMatrix.Get(row, col) < minVal {
				minVal = costMatrix.Get(row, col)
			}

		}
	}
	return minVal
}
