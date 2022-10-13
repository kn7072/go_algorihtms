package main

import (
	"fmt"
	"math"
)

func getHeightTable(lenArray int) int {
	return int(math.Floor(math.Log2(float64(lenArray))))
}

func setTableMin(sparseTable [][]int) {
	maxLevel := getHeightTable(len(sparseTable))
	lenRow := len(sparseTable[0])

	for level := 1; level <= maxLevel+1; level++ {
		for elemIndex := range sparseTable[level] {
			rightBorder := elemIndex + int(math.Pow(2, float64(level)))
			if rightBorder <= lenRow {
				sparseTable[level][elemIndex] = int(math.Max(float64(sparseTable[level-1][elemIndex]),
					float64(sparseTable[level-1][elemIndex+int(math.Pow(2, float64(level-1)))])))
			}
		}
	}
}

func request(sparseTable [][]int, left, right int) int {
	if left >= 0 && right >= 0 {
		if left == right {
			return sparseTable[left][0]
		} else {
			// todo : предподсчитать логарифмы для всех чисел < длины массива
			level := int(math.Floor(math.Log2(float64(right - left))))
			if level == 0 {
				return int(math.Max(float64(sparseTable[level][left]),
					float64(sparseTable[level][right])))
			} else {
				return int(math.Max(float64(sparseTable[level][left]),
					float64(sparseTable[level][right-int(math.Pow(2, float64(level)))])))
			}
		}
	} else {
		panic(fmt.Sprintf("Одна из границ отрицательная left %v, right %v", left, right))
	}
}

func createEmptyTable(array []int) [][]int {
	lenTable := len(array)
	heightTable := getHeightTable(lenTable) + 1
	sparseTable := make([][]int, heightTable)

	for i := 0; i < heightTable; i++ {
		sparseTable[i] = make([]int, lenTable)
	}

	copy(sparseTable[0], array)

	return sparseTable
}

func main() {
	array := []int{2, 4, 5, 8, 9, 1, 3, 7, 10}
	// fmt.Println(getHeightTable(len(array)))

	sparseTable := createEmptyTable(array)

	setTableMin(sparseTable)
	// fmt.Println(sparseTable)

	req := request(sparseTable, 1, 5)
	fmt.Println(req)

	req = request(sparseTable, 2, 3)
	fmt.Println(req)

	req = request(sparseTable, 0, 1)
	fmt.Println(req)

	req = request(sparseTable, 0, 0)
	fmt.Println(req)

	// req = request(sparseTable, -1, 0)
	// fmt.Println(req)
}
