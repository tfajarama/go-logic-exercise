package logic2

import (
	"github.com/tfajarama/go-logic-exercise/logic1"
	"github.com/tfajarama/go-logic-exercise/utils"
	"strconv"
)

// Number 1-2
func Logic2AscRows(num int, initial int, jump int) (result [][]int) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
	result = utils.InitMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			result[row][col] = numSlice[col] // fill with the value, forwards, same for each row
		}
	}

	return result
}

// Number 3
func Logic2AscAll(num int, initial int, jump int) (result [][]int) {
	value := initial
	result = utils.InitMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			result[row][col] = value // fill with the value, forwards, ascending each time
			value += jump
		}
	}

	return result
}

// Number 4
func Logic2AscAllCustomRow1Col1(num int, initial int, jumpRow1 int, jumpCol1 int, jumpOther int) (result [][]int) {
	value := initial
	result = utils.InitMatrix(num)

	for row := 0; row < num; row++ {
		if row == 0 {
			for col := 0; col < num; col++ {
				result[row][col] = value // fill with the value, forwards, ascending each time
				value += jumpRow1
			}
		} else {
			for col := 0; col < num; col++ {
				result[row][col] = value // fill with the value, forwards, ascending each time, with conditional case in edge columns
				if col == 0 || col == num-1 {
					value += jumpCol1
				} else {
					value += jumpOther
				}
			}
		}
	}

	return result
}

// Number 5
func Logic2AscAllAlternating(num int, initial int, jump int) (result [][]int) {
	value := initial
	result = utils.InitMatrix(num)

	for row := 0; row < num; row++ {
		if (row+1)%2 == 1 {
			for col := 0; col < num; col++ {
				result[row][col] = value // fill with the value, forwards, ascending each time
				value += jump
			}
		} else {
			for col := num - 1; col >= 0; col-- {
				result[row][col] = value // fill with the value, backwards, ascending each time
				value += jump
			}
		}
	}

	return result
}

// Number 6
func Logic2AscAllAlternatingCustomOddEvenRow(num int, initial int, jumpOddRow int, jumpEvenRow int, jumpBetweenRow int) (result [][]int) {
	value := initial
	result = utils.InitMatrix(num)

	for row := 0; row < num; row++ {
		if (row+1)%2 == 1 {
			for col := 0; col < num; col++ {
				result[row][col] = value // fill with the value, forwards, ascending each time
				if col == num-1 {
					value += jumpBetweenRow
				} else {
					value += jumpOddRow
				}
			}
		} else {
			for col := num - 1; col >= 0; col-- {
				result[row][col] = value // fill with the value, backwards, ascending each time
				if col == 0 {
					value += jumpBetweenRow
				} else {
					value += jumpEvenRow
				}
			}
		}
	}

	return result
}

// Number 7
func Logic2StairsDown(num int, initial int, jump int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if row == col { // if the row and col index is same (creating a stairs down), fill with the value
				result[row][col] = strconv.Itoa(numSlice[col])
			} else { // other than that, fill with an empty space
				result[row][col] = " "
			}
		}
	}

	return result
}

// Number 8
func Logic2StairsUp(num int, initial int, jump int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if row+col == num-1 { // if the row + col index is equals with the last index (creating a stairs up), fill with the value
				result[row][col] = strconv.Itoa(numSlice[col])
			} else { // other than that, fill with an empty space
				result[row][col] = " "
			}
		}
	}

	return result
}

// Number 9
func Logic2XSign(num int, initial int, jump int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if row+col == num-1 || row == col { // if the index creating a stairs down or a stairs up, fill with the value
				result[row][col] = strconv.Itoa(numSlice[col])
			} else { // other than that, fill with an empty space
				result[row][col] = " "
			}
		}
	}

	return result
}

// Number 10
func Logic2TriangleDownLeft(num int, initial int, jump int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if col <= row {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = " "
			}
		}
	}

	return result
}

// Number 11
func Logic2TriangleUpRight(num int, initial int, jump int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			if col >= row {
				result[row][col] = strconv.Itoa(numSlice[col])
			} else {
				result[row][col] = " "
			}
		}
	}

	return result
}

// Number 12
func Logic2TriangleLeftRight(num int, initial int, jump int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
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
				result[row][col] = " "
			}
		}
	}

	return result
}

// Number 13
func Logic2TriangleUpDown(num int, initial int, jump int) (result [][]string) {
	numSlice := logic1.Logic1AscStep(num, initial, jump)
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
				result[row][col] = " "
			}
		}
	}

	return result
}
