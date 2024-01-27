package main

import "fmt"

func main() {
	number := 123
	reversNum := numRevers(number)
	fmt.Println(number, " before revers")
	fmt.Println(reversNum, "  after revers")
}
func numRevers(n int) int {
	result := 0
	for n > 0 {
		remender := n % 10
		result = result*10 + remender
		n = n / 10
	}
	return result
}
