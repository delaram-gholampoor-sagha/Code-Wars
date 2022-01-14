package kata

func CountPositivesSumNegatives(numbers []int) []int {

	res := make([]int, 2)
	if len(numbers) == 0 || numbers == nil {
		return res
	}

	for _, v := range numbers {

		if v > 0 {

			res[0] += 1
		} else {
			res[1] += v
		}
	}

	return res // your code here
}
