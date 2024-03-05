package main

import "fmt"

func main() {
	num := 24
	result := listAllFactors(num)
	fmt.Println(result)
}
func listAllFactors(num int) []int {
	arr1 := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			arr1 = append(arr1, i)
		}
	}
	return arr1
}
