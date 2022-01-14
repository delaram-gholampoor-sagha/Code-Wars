package kata

func Grow(arr []int) int {
	result := 1
	for _, v := range arr {
		result *= v
	}
	return result
}
