package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123 45 67 89"
	wordsCount := strings.Split(str, " ")
	fmt.Println(len(wordsCount), " form keyword", wordsCount)
	result := sumOfCharacter(str)
	fmt.Println(result, " form function sum of character")
}
func sumOfCharacter(s string) int {
	count := 0
	for _, v := range s {
		if v == ' ' {
			continue
		} else {
			count++
		}
	}
	return count
}
