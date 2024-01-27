package main

import "fmt"

func main() {
	str := "abcd"
	fmt.Println(str, "  before permutation")
	resultstr := permutation(str)
	fmt.Println(resultstr, "   after permutation")
}
func permutation(s string) string {
	r := []rune(s)
	start, end := 0, len(s)-1
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
	return string(r)
}
