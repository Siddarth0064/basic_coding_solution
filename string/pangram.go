package main

import "fmt"

func main() {
	str1 := "abcdefghijklmnopqrstuvwxyz"
	resul1 := pengramChecking(str1)
	fmt.Println("the pangram is ", resul1, " from this string")
}
func pengramChecking(s1 string) bool {
	character := make(map[rune]bool)
	for _, v := range s1 {
		if v >= 'a' && v <= 'z' {
			character[v] = true
		}
	}
	return len(character) == 26
}
