package logic3

import (
	"github.com/tfajarama/go-logic-exercise/utils"
	"strconv"
)

// Number 1
func Logic3TriangleDownLeftAscAll(num int, initial int, jump int) (result [][]string) {
	value := initial
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		if (row+1)%2 == 1 {
			for col := 0; col < num; col++ {
				if col <= row {
					result[row][col] = strconv.Itoa(value)
					value += jump
				} else {
					result[row][col] = " "
				}
			}
		} else {
			for col := num - 1; col >= 0; col-- {
				if col <= row {
					result[row][col] = strconv.Itoa(value)
					value += jump
				} else {
					result[row][col] = " "
				}
			}
		}
	}

	return result
}

// Number 2
func Logic3TriangleUpRightAscAll(num int, initial int, jump int) (result [][]string) {
	value := initial
	result = utils.InitStrMatrix(num)

	for row := 0; row < num; row++ {
		if (row+1)%2 == 1 {
			for col := 0; col < num; col++ {
				if col >= row {
					result[row][col] = strconv.Itoa(value)
					value += jump
				} else {
					result[row][col] = " "
				}
			}
		} else {
			for col := num - 1; col >= 0; col-- {
				if col >= row {
					result[row][col] = strconv.Itoa(value)
					value += jump
				} else {
					result[row][col] = " "
				}
			}
		}
	}

	return result
}
