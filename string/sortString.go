package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "192837465"
	result := stringSort(str)
	fmt.Println(result)
}
func stringSort(s string) string {
	char := []rune(s)
	sort.Slice(char, func(i, j int) bool {
		return char[i] < char[j]
	})
	return string(char)
}
