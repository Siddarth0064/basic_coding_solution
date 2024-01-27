package main

import "fmt"

func divide(numerator, denominator int) {
	if denominator == 0 {
		panic("denominator is zero")
	}
	fmt.Println("no problem in denominator")
}
func safeDivied(numerator, denominator int) {
	defer recovery()
	divide(numerator, denominator)

}
func main() {
	fmt.Println("main starting")
	safeDivied(10, 5)
	safeDivied(10, 0)
	fmt.Println("main closing softly")
}

func recovery() {
	m := recover()

	if m != nil {
		fmt.Println(m)

	}
}
