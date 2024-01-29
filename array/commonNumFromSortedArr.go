package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 5, 6, 7, 45, 48}
	arr2 := []int{2, 3, 5, 7, 48}
	resultArr := intersection(arr1, arr2)
	fmt.Println(resultArr)
}
func intersection(a1, a2 []int) []int {
	resultArr := []int{}
	start1 := 0
	start2 := 0
	for start1 < len(a1) && start2 < len(a1) {
		if a1[start1] == a2[start2] {
			resultArr = append(resultArr, a1[start1])
			start1++
			start2++
		} else if a1[start1] > a2[start2] {
			start2++
		} else {
			start1++
		}

	}
	return resultArr

}
