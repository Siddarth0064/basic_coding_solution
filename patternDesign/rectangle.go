package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j < 5; j++ {
			if i == 1 || i == 5 || j == 1 || j == 5 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
