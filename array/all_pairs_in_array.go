package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	resArray := [][]int{}
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			resArray = append(resArray, []int{arr1[i], arr1[j]})
		}
	}
	fmt.Println(resArray)
}
