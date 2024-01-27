package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chan1 := make(chan int)
	chan2 := make(chan int, 10)
	chan3 := make(chan int)
	go SumofNumber(slice1, chan1)
	go evenNumber(slice1, chan2)
	go heigstNum(slice1, chan3)
	result := <-chan1
	fmt.Println(result, " is  the number of sum in the slice")

	for evenResult := range chan2 {
		fmt.Println(evenResult)
	}

	maxNum := <-chan3
	fmt.Println(maxNum, "this is the maximum number in the slice")

}
func SumofNumber(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	close(c)
}
func evenNumber(num []int, ch chan int) {
	for i := 0; i < len(num); i++ {
		if num[i]%2 == 0 {
			ch <- num[i]
		}
	}
	close(ch)
}
func heigstNum(s []int, ch chan int) {
	max := 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	ch <- max
}
func miniNum(s []int)
