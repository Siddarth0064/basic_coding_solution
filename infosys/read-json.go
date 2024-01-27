package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	jsonbyte, err := os.ReadFile("jons-file.txt")
	if err != nil {
		fmt.Println("there is an err in readfile", err)
	}
	fmt.Println(string(jsonbyte))
	//stringData := string(jsonbyte)
	// linesSpliteFile := strings.Split(stringData, "\n")
	// fmt.Println(linesSpliteFile)
	// for linesData, _ := range stringData {
	// 	lData := strings.Split(linesData, ",")
	// 	fmt.Println(lData)
	// }
	linesData := strings.Split(string(jsonbyte), ",")
	fmt.Println(linesData)
	for _, details := range linesData {
		//	dOfl := strings.Split(details, "\n")
		fmt.Println(details)
	}
	number := 2
	num, err := strconv.Itoa(number)
	if err != nil {
		fmt.Println("its an error of strconvr")
	}
	fmt.Prinln(num)
}
