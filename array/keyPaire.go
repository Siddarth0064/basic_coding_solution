package main

import "fmt"

func main() {
	num := 16
	arr1 := []int{1, 4, 5, 6, 10, 8}
	reultArr, resultMap := pairChecking(arr1, num)
	fmt.Println(reultArr, resultMap)
}
func pairChecking(a []int, n int) (bool, map[int]int) {
	mapContain := make(map[int]int)
	for i := 0; i < len(a); i++ {
		mapContain[a[i]]++
		val := n - a[i]
		_, ok := mapContain[val]
		if ok {
			return true, mapContain
		}

	}
	return false, mapContain
}
