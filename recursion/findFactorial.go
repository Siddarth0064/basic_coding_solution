package main

import "fmt"

func main() {
	result := fact(10)
	fmt.Println(result)
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) * n
}
