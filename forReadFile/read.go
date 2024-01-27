package main

import (
	"fmt"
	"os"
	//"strings"
	//
	//"io/ioutil"
)

func main() {
	arr, err := os.ReadFile("abcd.txt")
	if err != nil {
		fmt.Println("file is not found")
	}
	fmt.Println(string(arr))
	str := string(arr)

	fmt.Println(str)
	// name := []byte("hi im writeing in vs code")
	fi, err := os.Create("aaa.txt")
	if err != nil {
		fmt.Println("un able to create the file")
	}
	defer fi.Close()
	fileabc, err := os.Create("abc.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string())
	fiel11, err := os.ReadFile("aaa.txt")
	srt1 := string(fiel11)
	fmt.Println(srt1)
	countW := 0
	countS := 0
	for _, v := range srt1 {
		if v == ' ' {
			countW++
			countS++
		}
	}
	fmt.Println("the number of words is :", countW)
	fmt.Println("the number of letters is: ", len(srt1)-countS)
	// err11 := os.Remove("aaa.txt")
	// fmt.Println(err11)
	op, err := os.Open("aa.txt")
	fmt.Println(op)

}
