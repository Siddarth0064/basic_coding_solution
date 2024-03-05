package main

import "fmt"

func main() {
	result := findYcommonMultiples(3, 5, 4)
	fmt.Println(result)
}
func findYcommonMultiples(n1, n2, y int) []int {
	mul := n1 * n2
	arr1 := []int{}
	for y != 0 {
		arr1 = append(arr1, mul)
		mul += n1 * n2
		y--
	}
	return arr1
}
