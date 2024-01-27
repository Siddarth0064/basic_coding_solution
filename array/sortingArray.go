package main

import (
	"fmt"
	"math/rand"
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
