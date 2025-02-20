package logic1

import "strconv"

// Number 1 - 3
func Logic1AscStep(num int, initial int, jump int) (result []int) {
	result = make([]int, num)
	value := initial

	for i := 0; i < num; i++ { // fill the value, forwards
		result[i] = value
		value += jump
	}

	return result
}

// Number 4 - 6
func Logic1DescStep(num int, initial int, jump int) (result []int) {
	result = make([]int, num)
	value := initial

	for i := num - 1; i >= 0; i-- { // fill with the value, backwards
		result[i] = value
		value += jump
	}

	return result
}

// Number 7 - 9
func Logic1AscDescStep(num int, initial int, jump int) (result []int) {
	result = make([]int, num)
	value := initial
	midPoint := initial
	isEven := false

	// check if num series is even or odd
	if num%2 == 0 {
		midPoint = (num / 2) - 1
		isEven = true
	} else {
		midPoint = ((num + 1) / 2) - 1
	}

	for i := 0; i < num; i++ {
		result[i] = value
		if i == midPoint && isEven { // if the index is in the midpoint and the series is even, do nothing to the value
			continue
		} else if i >= midPoint { // if the index is equal (only if odd) or more than the midpoint, minus the value
			value -= jump
		} else { // if the index is before the midpoint, add the value
			value += jump
		}
	}

	return result
}

// Number 10
func Logic1OddNumEvenFizz(num int, initial int, jump int) (result []string) {
	result = make([]string, num)
	value := initial

	for i := 0; i < num; i++ {
		if (i+1)%2 == 0 { // if index is even, fill with "fizz"
			result[i] = "fizz"
		} else { // if index is odd, fill with the value
			result[i] = strconv.Itoa(value)
			value *= jump
		}
	}

	return result
}

// Number 11
func Logic1OddBuzzEvenNum(num int, initial int, jump int) (result []string) {
	result = make([]string, num)
	value := initial

	for i := 0; i < num; i++ {
		if i == 1 { // if the index is 1, fill with plain "1"
			result[i] = "1"
		} else if (i+1)%2 == 1 { // if the index is odd, fill with "buzz"
			result[i] = "buzz"
		} else { // if the index is even (except the '1' index), fill with the value
			result[i] = strconv.Itoa(value)
			value *= jump
		}
	}

	return result
}

// Number 12
func Logic1Repeat4Elements(num int, initial int, jump int) (result []int) {
	numSlice := Logic1AscStep(4, initial, jump)
	result = make([]int, num)
	pointer := 0

	for i := 0; i < num; i++ {
		result[i] = numSlice[pointer]
		if (i+1)%4 == 0 { // if the index is divisible by 4, reset the value pointer
			pointer = 0
		} else { // if not, add 1 to the value pointer
			pointer++
		}
	}

	return result
}
