package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%v", j)
		}

		for j := 1; j <= i; j++ {
			fmt.Printf("%v", j)
		}
		for j := 5; j >= i; j-- {
			fmt.Print(" ")
		}
		for j := 5; j >= i; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%v", j)
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%v", j)
		}
		fmt.Println()

	}

}
