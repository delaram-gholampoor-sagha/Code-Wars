package main

import "fmt"

func main() {
	fmt.Println(Feast(" delaram ", "vegetables"))

}

func Feast(beast string, dish string) bool {
	return beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1]
}
