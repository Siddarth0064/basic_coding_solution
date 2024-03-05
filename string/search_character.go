package main

import (
	"fmt"
)

func main() {
	words := []string{"aaaaa", "accbbaaaaa", "aabbaaaaa", "abbb", "addd", "accccc", "abc"}
	rsult := sortCharacters(words)
	fmt.Println(rsult)
}
func sortCharacters(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words)-i-1; j++ {
			if words[j] > words[j+1] {
				words[j], words[j+1] = words[j+1], words[j]
			}
		}
	}
	return words
}
