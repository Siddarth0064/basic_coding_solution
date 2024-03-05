package main

import "fmt"

func main() {
	num := 12321
	tempNum := num
	var reversNum int
	for tempNum != 0 {
		remender := tempNum % 10
		reversNum = reversNum*10 + remender
		tempNum /= 10
	}
	isPalindrome := num == reversNum
	fmt.Println(isPalindrome)
}
