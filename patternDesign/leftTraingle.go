package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for s := 1; s <= 5-i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*" + " ")
		}

		fmt.Println()
	}
}
