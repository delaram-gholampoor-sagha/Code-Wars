package main

import "fmt"

func main() {

	fmt.Println(Tickets([]int{25}))
	fmt.Println(Tickets([]int{25, 100}))
	fmt.Println(Tickets([]int{25, 25, 50, 50, 100}))
}

func Tickets(peopleInLine []int) string {
	var (
		money_25  int
		money_50  int
		money_100 int
	)
	for _, v := range peopleInLine {
		if v == 25 {
			money_25++
		} else if v == 50 && money_25 > 0 {
			money_50++
			money_25--
		} else if v == 100 && money_25 > 0 && money_50 > 0 {
			money_100++
			money_50--
			money_25--
		} else if v == 100 && money_25 > 2 {
			money_100++
			money_25 -= 3
		} else {
			return "NO"
		}

	}

	return "YES"

}
