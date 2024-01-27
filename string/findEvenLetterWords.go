package main

import "fmt"

func main() {
	str := "sky is blue and vast"
	resultStr := findEvenWords(str)
	fmt.Println(resultStr)
}
func findEvenWords(s string) []string {
	wordsmap := make(map[string]int)
	slice1 := []string{}
	st := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			word := s[st:i]
			wordsmap[word] = len(word)
			if wordsmap[word]%2 == 0 {
				slice1 = append(slice1, word)
			}
			st = i + 1
		}

	}
	return slice1
}
