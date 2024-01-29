package main

import "fmt"

func main() {
	str := "eaabbaa"
	palindromeResult, ok := palindrome(str)
	fmt.Println(palindromeResult, "    ", ok, " ", str)

}
func palindrome(s string) (string, bool) {
	restStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		restStr += string(s[i])
	}
	if s == restStr {
		return restStr, true
	}
	return restStr, false
}
