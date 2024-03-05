package main

import "fmt"

func main() {
	num := 6
	result := findPrimeOrNot(num)
	fmt.Println(result)

}
func findPrimeOrNot(num int) string {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return "Not Prime"
		}
	}
	return "its Prime"
}
