package main

import "fmt"

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 0, 6, 0}
	fmt.Println(arr1, " before moving zero's ")
	resultArr := moveZerosToLast(arr1)
	fmt.Println(resultArr, " after moving zero's ")

}
func moveZerosToLast(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == 0 {
			for j := i + 1; j < len(a); j++ {
				if a[j] != 0 {
					a[i], a[j] = a[j], a[i]
					break
				}
			}
		}
	}
	return a
}
