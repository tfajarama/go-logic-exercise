package logic1

import "strconv"

// Number 1 - 3
func Logic1AscStep(num int, initial int, step int) (slice []int) {
	slice = make([]int, num)
	numPoint := initial
	for i := 0; i < num; i++ {
		slice[i] = numPoint
		numPoint += step
	}
	return slice
}

// Number 4 - 6
func Logic1DescStep(num int, initial int, step int) (slice []int) {
	slice = make([]int, num)
	numPoint := initial
	for i := num - 1; i >= 0; i-- {
		slice[i] = numPoint
		numPoint += step
	}
	return slice
}

// Number 7 - 9
func Logic1AscDescStep(num int, initial int, step int) (slice []int) {
	slice = make([]int, num)
	numPoint := initial
	midPoint := initial
	isEven := false

	if num%2 == 0 {
		midPoint = num / 2
		isEven = true
	} else {
		midPoint = (num + 1) / 2
	}

	for i := 0; i < num; i++ {
		slice[i] = numPoint
		if i == midPoint-1 && isEven {
			continue
		} else if i >= (midPoint - 1) {
			numPoint -= step
		} else {
			numPoint += step
		}
	}
	return slice
}

// Number 10
func Logic1OddNumEvenFizz(num int, initial int, step int) (slice []string) {
	slice = make([]string, num)
	numPoint := initial
	for i := 0; i < num; i++ {
		if (i+1)%2 == 0 {
			slice[i] = "fizz"
		} else {
			slice[i] = strconv.Itoa(numPoint)
			numPoint *= step
		}
	}
	return slice
}

// Number 11
func Logic1OddBuzzEvenNum(num int, initial int, step int) (slice []string) {
	slice = make([]string, num)
	numPoint := initial
	for i := 0; i < num; i++ {
		if i == 1 {
			slice[i] = "1"
		} else if (i+1)%2 == 1 {
			slice[i] = "buzz"
		} else {
			slice[i] = strconv.Itoa(numPoint)
			numPoint *= step
		}
	}
	return slice
}
