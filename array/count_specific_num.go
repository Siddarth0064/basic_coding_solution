package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 2, 2, 4, 5}
	elementTOcount := 2
	var occurren int
	for _, val := range arr1 {
		if elementTOcount == val {
			occurren++
		}
	}
	fmt.Println(occurren)
}
