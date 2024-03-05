package main

import "fmt"

func main() {
	arr1 := []int{1, 1, 1, 2, 3, 3, 5, 5, 5}
	var maxRepeating int
	maxOccurren := 0
	for i := 0; i < len(arr1); i++ {
		count := 0
		for j := 0; j < len(arr1); j++ {
			if arr1[i] == arr1[j] && i != j {
				count++
			}

		}
		if count >= maxOccurren {
			maxOccurren = count
			maxRepeating = arr1[i]
		}
	}
	fmt.Println(maxRepeating)
}
