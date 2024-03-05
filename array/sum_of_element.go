package main

import "fmt"

func main() {
	sum := 0
	arr1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println(sum)
}
