package kata

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func PositiveSum(numbers []int) int {
	result := 0
	for i, v := range numbers {
		if v < 0 {
			remove(numbers, i)
		} else {
			result += v
		}
	}
	return result
}
