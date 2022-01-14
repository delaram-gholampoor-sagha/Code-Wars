package kata

func SquareSum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v * v
	}
	return result
}
