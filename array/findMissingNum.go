package main

import "fmt"

func main() {
	arr1 := []int{1, 4, 2, 9, 5, 6}
	result := findMissingNum(arr1)
	fmt.Println(result, "  these are the missing numbers from slice")
}
func findMissingNum(a []int) []int {
	slice1 := []int{}
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]

			}
		}
	}
	min := a[0]
	max := a[len(a)-1]
	numbermap := make(map[int]bool)
	for _, v := range a {
		if !numbermap[v] {
			numbermap[v] = true
		}
	}
	fmt.Println(numbermap, min, max, a)

	for i := min; i <= max; i++ {
		if numbermap[i] == false {
			slice1 = append(slice1, i)
		}
	}
	return slice1
}
