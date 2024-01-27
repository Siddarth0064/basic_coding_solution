package main //// sum of even form 1 to 100

import "fmt"

func main() {
	slice1 := []int{}
	for i := 0; i < 101; i++ {
		slice1 = append(slice1, i)
	}
	ch := make(chan int)
	go sumOfEven(slice1, ch)
	result := <-ch
	fmt.Println(result)
}
func sumOfEven(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		if v%2 == 0 {
			sum += v
		}
	}
	ch <- sum
	close(ch)
}
