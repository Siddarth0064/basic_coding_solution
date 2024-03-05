package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{5, 10, 3, 8, 15}
	var largestNum int
	for _, val := range arr1 { // this is one way to find largest num
		if val > largestNum {
			largestNum = val
		}
	}
	fmt.Println(largestNum)
	findLargestNum(arr1)
}

func findLargestNum(arr1 []int) { // this is another way to find largest num
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	fmt.Println(arr1[len(arr1)-1])
}
