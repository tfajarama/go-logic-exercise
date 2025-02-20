package utils

func InitMatrix(num int) (result [][]int) {
	result = make([][]int, num)
	for i := range result {
		result[i] = make([]int, num)
	}
	return result
}

func InitStrMatrix(num int) (result [][]string) {
	result = make([][]string, num)
	for i := range result {
		result[i] = make([]string, num)
	}
	return result
}
