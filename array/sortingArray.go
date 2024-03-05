package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	array := []int{}
	for i := 0; i < 10; i++ {
		array = append(array, rand.Intn(100))
	}
	fmt.Println(array)
	result := sorting(array)
	fmt.Println(result)
	arr2 := []int{1, 3, 2, 4, 5}
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	fmt.Println(arr2, " by  using anonymous function")
}
func sorting(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
