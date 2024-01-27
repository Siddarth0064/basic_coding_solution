package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr1 := []int{}
	for i := 0; i < 100; i++ {
		arr1 = append(arr1, src.Intn(100)+100)
	}
	fmt.Println(arr1)
	ch := make(chan int, 2)
	go FindMinMax(arr1, ch)
	for result := range ch {
		fmt.Println(result)
	}
}
func FindMinMax(s []int, ch chan int) {
	max, min := s[0], s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	ch <- max
	ch <- min
	close(ch)
}
