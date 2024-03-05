package main

import "fmt"

func main() {
	n1 := 12
	n2 := 18
	result := findCommonFactors(n1, n2)
	fmt.Println(result)
}

func findCommonFactors(x, y int) []int {
	minNum := min(x, y)
	arr1 := []int{}
	for i := 1; i <= minNum; i++ {
		if x%i == 0 && y%i == 0 {
			arr1 = append(arr1, i)
		}
	}
	return arr1
}
