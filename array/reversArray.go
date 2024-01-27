package main

import "fmt"

//import "math/rand"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr1, "  befor revering")
	resulArray := reversArray(arr1)
	fmt.Println(resulArray, "  after revering")
}
func reversArray(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
