package kata

import "strings"

// this function is accepted in codewars solution .
func RepeatStr(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}

// this is my own function it produces the exact same result but it wasn't accepted in solutions .
// func RepeatStr(repetitions int, value string) string {

// 	for i := 1; i < repetitions; i++ {
// 		fmt.Print(value)
// 	}
// 	return value
// }
