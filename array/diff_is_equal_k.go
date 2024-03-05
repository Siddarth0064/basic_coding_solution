package main

import "fmt"

func main() {
	arr1 := []int{9, 5, 6, 1, 2}
	k := 4
	re := [][]int{}
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i]-arr1[j] == k {
				re = append(re, []int{arr1[i], arr1[j]})
			}
		}

	}
	fmt.Println(re)
}
