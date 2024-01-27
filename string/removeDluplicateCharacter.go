package main

import "fmt"

func main() {
	str := "siddarth"
	finalStr := removeDluplicate(str)
	fmt.Println(finalStr)
}
func removeDluplicate(s string) string {
	strmap := make(map[rune]bool)
	slice := []rune{}
	for _, v := range s {
		if !strmap[v] {
			strmap[v] = true
			slice = append(slice, v)
		}
	}
	return string(slice)
}
