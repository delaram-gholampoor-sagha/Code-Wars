package multiplicationtable

func MultiplicationTable(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		for x := 1; x < size+1; x++ {
			res[i] = append(res[i], (i+1)*x)
		}
	}
	return res

}
