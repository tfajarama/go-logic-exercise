package logic3

import (
	"github.com/tfajarama/go-logic-exercise/utils"
)

// Number 1
func Number1(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	value := 1
	for row := 0; row < num; row++ {
		for col := 0; col <= row; col++ {
			if row%2 == 0 {
				result[row][col] = value
			} else {
				result[row][row-col] = value
			}
			value += 2
		}
	}
	return result
}

// Number 2
func Number2(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	value := 1
	for row := 0; row < num; row++ {
		for col := row; col < num; col++ {
			if row%2 == 0 {
				result[row][col] = value
			} else {
				result[row][(num-1)+(row-col)] = value
			}
			value += 2
		}
	}
	return result
}

// Number 3 : UPDATE LATER
func Number3(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	value := 2
	for row := 0; row < num; row++ {
		for col := 0; col < num-row; col++ {
			if row%2 == 0 {
				result[row][col] = value
				value += 2
			} else {
				result[row][num-1-col-row] = value
				value += 3
			}
		}
	}
	return result
}

// Number 4
func Number4(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	value := 1
	for row := 0; row < num; row++ {
		for col := num - 1 - row; col < num; col++ {
			if row%2 == 1 {
				result[row][col] = value
			} else {
				result[row][(num-1-col)+(num-1-row)] = value
			}
			value += 2
		}
	}
	return result
}

// Number 5
func Number5(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num + 1) / 2
	value := 1
	// process only a quarter (top-left), mirror the other 3 quarter
	for row := 0; row < midNum; row++ {
		// iterate only half rows, and iterate cols inside the left & right edges
		for col := 0; col <= row; col++ {
			if row%2 == 0 { // if the row index is even, fill from the left
				result[row][col] = value             // top-left quarter
				result[row][num-1-col] = value       // mirror top-right quarter
				result[num-1-row][col] = value       // mirror bottom-left quarter
				result[num-1-row][num-1-col] = value // mirror bottom-right quarter
			} else { // if the row index is odd, fill from the right
				result[row][row-col] = value               // top-left quarter
				result[row][col+(num-1-row)] = value       // mirror top-right quarter
				result[num-1-row][row-col] = value         // mirror bottom-left quarter
				result[num-1-row][col+(num-1-row)] = value // mirror bottom-right quarter
			}
			value += 2
		}
	}
	return result
}

// Number 6
func Number6(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num + 1) / 2
	value := 1
	// iterate only the top half rows, mirror the bottom half
	for row := 0; row < midNum; row++ {
		// 	iterate inside the left & right edges
		for col := row; col < num-row; col++ {
			if row%2 == 0 { // if the row index is even, fill from the left
				result[row][col] = value       // top half
				result[num-1-row][col] = value // mirror bottom half
			} else { // if the row index is odd, fill from the right
				result[row][num-1-col] = value       // top half
				result[num-1-row][num-1-col] = value // mirror bottom half
			}
			value += 2
		}
	}
	return result
}

// Number 7
func Number7(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num - 1) / 2
	value := 1
	// iterate only the top half rows, mirror the bottom half
	for row := 0; row < (num+1)/2; row++ {
		// 	iterate inside the left & right edges
		for col := midNum - row; col <= midNum+row; col++ {
			if row%2 == 1 { // if the row index is odd, fill from the left
				result[row][col] = value       // top half
				result[num-1-row][col] = value // mirror bottom half
			} else { // if the row index is even, fill from the right
				result[row][num-1-col] = value       // top half
				result[num-1-row][num-1-col] = value // mirror bottom half
			}
			value += 2
		}
	}
	return result
}

// Number 8
func Number8(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num - 1) / 2
	value := 1
	// iterate only the left half columns, mirror the right half
	for col := 0; col < (num+1)/2; col++ {
		// iterate inside the top & bottom edges
		for row := midNum - col; row <= midNum+col; row++ {
			if col%2 == 1 { // if the col index is odd, fill from the top
				result[row][col] = value       // left half
				result[row][num-1-col] = value // mirror right half
			} else { // if the col index is even, fill from the bottom
				result[num-1-row][col] = value       // left half
				result[num-1-row][num-1-col] = value // mirror right half
			}
			value += 2

		}
	}
	return result
}

// Number 9
func Number9(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num - 1) / 2
	for row := midNum; row >= 0; row-- {
		value := 1
		for col := midNum - row; col <= midNum; col++ {
			result[row][col] = value
			result[row][num-1-col] = value
			result[num-1-row][col] = value
			result[num-1-row][num-1-col] = value
			value += 2
		}
	}
	return result
}

// Number 10
func Number10(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num - 1) / 2
	for row := midNum; row >= 0; row-- {
		value := 1 + (2 * (midNum - row))
		for col := midNum; col >= midNum-row; col-- {
			result[row][col] = value
			result[row][num-1-col] = value
			result[num-1-row][col] = value
			result[num-1-row][num-1-col] = value
			value += 2
		}
	}
	return result
}

// Number 11
func Number11(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num - 1) / 2
	valueLeftMid := 1  // store value for left half
	valueRightMid := 1 // store value for right half
	// iterate only the top half
	for row := 0; row <= midNum; row++ {
		// fill the left half
		result[row][row] = valueLeftMid       // top half
		result[num-1-row][row] = valueLeftMid // mirror bottom half
		valueLeftMid += 2
		// fill the right half
		for col := num - 1 - row; col < num; col++ {
			if row%2 == 1 { // if the row index is odd, fill from the left
				result[row][col] = valueRightMid       // top half
				result[num-1-row][col] = valueRightMid // mirror bottom half
			} else { // if the row index is even, fill from the right
				result[row][(num-1-col)+(num-1-row)] = valueRightMid       // top half
				result[num-1-row][(num-1-col)+(num-1-row)] = valueRightMid // mirror bottom half
			}
			valueRightMid += 2
		}
	}
	return result
}

// Number 11
func Number11b(num int) (result [][]int) {
	midNum := (num - 1) / 2
	result = utils.InitRectangleMatrix(num, num+midNum)
	valueLeftMid := 1  // store value for left half
	valueRightMid := 1 // store value for right half
	// iterate only the top half
	for row := 0; row <= midNum; row++ {
		// fill the left half
		result[row][row] = valueLeftMid       // top half
		result[num-1-row][row] = valueLeftMid // mirror bottom half
		valueLeftMid += 2
		// fill the right half
		for col := num - 1 - row; col < num; col++ {
			if row%2 == 1 { // if the row index is odd, fill from the left
				result[row][col] = valueRightMid                       // top half left quarter
				result[num-1-row][col] = valueRightMid                 // mirror bottom half left quarter
				result[row][(num-1-col)+(num-1)] = valueRightMid       // mirror top half right quarter
				result[num-1-row][(num-1-col)+(num-1)] = valueRightMid // mirror bottom half right quarter
			} else { // if the row index is even, fill from the right
				result[row][(num-1-col)+(num-1-row)] = valueRightMid       // top half left quarter
				result[num-1-row][(num-1-col)+(num-1-row)] = valueRightMid // mirror bottom half left quarter
				result[row][col+row] = valueRightMid                       // mirror top half right quarter
				result[num-1-row][col+row] = valueRightMid                 // mirror bottom half right quarter
			}
			valueRightMid += 2
		}
	}
	return result
}

// Number 12 Smart-Work Style
func Number12SWStyle(num int) (result [][]int) {
	mirrorResult := Number11(num)
	result = utils.InitMatrix(num)
	for row := 0; row < num; row++ {
		for col := 0; col < num; col++ {
			result[row][num-1-col] = mirrorResult[row][col]
		}
	}
	return result
}

// Number 12b Smart-Work Style
func Number12bSWStyle(num int) (result [][]int) {
	mirrorResult := Number11b(num)
	midNum := (num - 1) / 2
	result = utils.InitRectangleMatrix(num, num+midNum)
	for row := 0; row < len(mirrorResult); row++ {
		for col := 0; col < len(mirrorResult[row]); col++ {
			result[row][num+midNum-1-col] = mirrorResult[row][col]
		}
	}
	return result
}

// Number 12 Hard-Work Style
func Number12HWStyle(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	midNum := (num - 1) / 2
	valueLeftMid := 1  // store value for left half
	valueRightMid := 1 // store value for right half
	// iterate only the top half
	for row := 0; row <= midNum; row++ {
		// fill the right half
		result[row][num-1-row] = valueRightMid       // top half
		result[num-1-row][num-1-row] = valueRightMid // mirror bottom half
		valueRightMid += 2
		// fill the left half
		for col := 0; col <= row; col++ {
			if row%2 == 0 { // if the row index is even, fill from the left
				result[row][col] = valueLeftMid       // top half
				result[num-1-row][col] = valueLeftMid // mirror bottom half
			} else { // if the row index is odd, fill from the left
				result[row][row-col] = valueLeftMid       // top half
				result[num-1-row][row-col] = valueLeftMid // mirror bottom half
			}
			valueLeftMid += 2
		}
	}
	return result
}

// Number 12 Hard-Work Style
func Number12bHWStyle(num int) (result [][]int) {
	midNum := (num - 1) / 2
	result = utils.InitRectangleMatrix(num, num+midNum)
	valueLeftMid := 1  // store value for left half
	valueRightMid := 1 // store value for right half
	// iterate only the top half
	for row := 0; row <= midNum; row++ {
		// fill the right half
		result[row][num+midNum-1-row] = valueRightMid       // top half
		result[num-1-row][num+midNum-1-row] = valueRightMid // mirror bottom half
		valueRightMid += 2
		// fill the left half
		for col := midNum; col <= midNum+row; col++ {
			if row%2 == 0 { // if the row index is even, fill from the left
				result[row][col] = valueLeftMid             // top half right quarter
				result[num-1-row][col] = valueLeftMid       // mirror bottom half right quarter
				result[row][num-1-col] = valueLeftMid       // mirror top half left quarter
				result[num-1-row][num-1-col] = valueLeftMid // mirror bottom half left quarter
			} else { // if the row index is odd, fill from the left
				result[row][(row-col)+(num-1)] = valueLeftMid       // top half right quarter
				result[num-1-row][(row-col)+(num-1)] = valueLeftMid // mirror bottom half right quarter
				result[row][col-row] = valueLeftMid                 // mirror top half left quarter
				result[num-1-row][col-row] = valueLeftMid           // mirror bottom half left quarter
			}
			valueLeftMid += 2
		}
	}
	return result
}

// Number 14 : UPDATE LATER
func Number14(num int) (result [][]int) {
	result = utils.InitMatrix(num)
	value := 1
	// iterate the columns, then its rows
	for col := 0; col < num; col++ {
		for row := 0; row < num; row++ {
			if col%2 == 0 { // if the row index is even, fill from the top
				result[row][col] = value
			} else { // if the row index is odd, fill from the bottom
				result[num-1-row][col] = value
			}
			value += 2
		}
	}
	return result
}
