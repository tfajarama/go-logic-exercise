package logic2

import (
	"github.com/tfajarama/go-logic-exercise/logic1"
	"github.com/tfajarama/go-logic-exercise/utils"
	"strconv"
)

// Number 3
func Logic2AscAll(num int, initial int, step int) (result [][]int) {
	numPoint := initial
	result = utils.InitMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			result[row][col] = numPoint
			numPoint += step
		}
	}
	return result
}

// Number 7
func Logic2StairsDown(num int, initial int, step int) (result [][]string) {
	numPoint := initial
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if row == col {
				result[row][col] = strconv.Itoa(numPoint)
				numPoint += step
			} else {
				result[row][col] = "-"
			}
		}
	}
	return result
}

// Number 8
func Logic2StairsUp(num int, initial int, step int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, step)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if row+col == num-1 {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = "-"
			}
		}
	}
	return result
}

// Number 9
func Logic2XSign(num int, initial int, step int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, step)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if row+col == num-1 || row == col {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = "-"
			}
		}
	}
	return result
}

// Number 10
func Logic2TriangleDownLeft(num int, initial int, step int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, step)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if col <= row {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = "-"
			}
		}
	}
	return result
}

// Number 11
func Logic2TriangleUpRight(num int, initial int, step int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, step)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if col >= row {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = "-"
			}
		}
	}
	return result
}

// Number 12
func Logic2TriangleLeftRight(num int, initial int, step int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, step)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			head := row
			tail := num - 1 - row

			// if the head pointer is already bigger than the tail pointer,
			// then swap the value references
			if head > tail {
				temp := head
				head = tail
				tail = temp
			}

			if col <= head || col >= tail {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = "-"
			}
		}
	}
	return result
}

// Number 13
func Logic2TriangleUpDown(num int, initial int, step int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, step)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			head := row
			tail := num - 1 - row

			// if the head pointer is already bigger than the tail pointer,
			// then swap the value references
			if head > tail {
				temp := head
				head = tail
				tail = temp
			}

			if col >= head && col <= tail {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = "-"
			}
		}
	}
	return result
}
