package main

import "fmt"

func main() {
	//=============== slice of integer revers ==========================
	a := []int{1, 2, 3, 4, 5, 6}
	i := 0
	j := len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	fmt.Println(a)
	//============= slice of string revers =========================
	s := []string{"a", "b", "c", "d"}
	is := 0
	js := len(s) - 1
	for is < js {
		s[is], s[js] = s[js], s[is]
		is++
		js--
	}
	fmt.Println(s)
	//=============== string revers ==================================
	str := "siddarth"
	runes := []rune(str)
	for ii, jj := 0, len(runes)-1; ii < jj; ii, jj = ii+1, jj-1 {
		runes[ii], runes[jj] = runes[jj], runes[ii]
	}
	fmt.Println("before revers: ", str)
	fmt.Println("after the revers :", string(runes))
	//============= printing alternative charater from two string ===============
	str1 := "11111"
	str2 := "22222"
	var str3 string
	stri, strj := 0, 0
	for stri < len(str1) && strj < len(str2) {
		str3 += string(str1[stri]) + string(str2[strj])
		stri++
		strj++
	}
	str3 += str1[stri:]
	str3 += str2[strj:]
	fmt.Println("printing the alternative character from two string := ", str3)
	//==================== string palindrom check =================
	pstr := "apapa"
	for i, j := 0, len(pstr)-1; i < j; i, j = i+1, j-1 {
		if pstr[i] != pstr[j] {
			fmt.Printf("the given string is not  a palindrom  : %s\n", pstr)
			return
		}
	}
	fmt.Printf("the given string is a  palindrom : %s\n", pstr)
	//============= integer palindrom check =======================
	num := 10000
	finalnum := num
	fmt.Println("before the revers number is : ", num)
	var revers int
	for num != 0 {
		remender := num % 10
		revers = (revers * 10) + remender
		num = num / 10
	}
	fmt.Println("after the revers number is : ", revers)
	if finalnum == revers {
		fmt.Printf("the given number is  %v and %v is palindrom \n", finalnum, revers)
	} else {
		fmt.Printf("the given number is  %v and %v is not palindrom \n", finalnum, revers)
	}
	//============= remove element by crating another array  ================================
	rarr := []int{1, 2, 3, 4, 5, 6, 7}
	eletoremove := 4
	var resultarr []int
	for _, v := range rarr {
		if eletoremove != v {
			resultarr = append(resultarr, v)
		}
	}
	fmt.Println("before removed element from the arr ", rarr)
	fmt.Println("after the removed particuler element from thr arr ", resultarr)
	//================= with out creating another copy of an arr removing element ==============

	// numbers := []int{10, 10, 20, 30, 20, 10} //slice containes duplicate elements
	// fmt.Println("before removing the duplicate value : ", numbers)
	// count := 1
	// previous := numbers[0]
	// for i := 0; i < len(numbers); i++ {
	// 	if numbers[i] != previous {
	// 		numbers[count] = numbers[i]
	// 		count++
	// 		previous = numbers[i]
	// 	}
	// }
	// fmt.Println(count)

	// fmt.Println("before remove duplicate values: ", numbers)
	//===============duplicateRemove(numbers)

	// for i := 0; i < len(rarr); i++ {
	// 	if eletoremove == rarr[i] {
	// 		copy(rarr[i:], rarr[i+1:])
	// 		rarr = rarr[:len(rarr)-1]
	// 		i--
	// 	}
	// }
	// fmt.Println("reming element form the original arr : ", rarr)

	//============= ficonaic series of elemrnt ==================================================
	nums := 6
	initial := 0
	first := 1
	if nums == 0 {
		fmt.Println(initial)
	} else if nums == 1 {
		fmt.Println(initial, " ", first)
	} else {
		fmt.Print(initial, " ", first)
		for i := 0; i < nums; i++ {
			thired := initial + first
			fmt.Print(" ", thired)
			initial = first
			first = thired

		}

	}
	//======================= muximum sum of sub array in the array =========================
	darr := []int{1, 2, -5, 4, 5, 1, 2, 3, 4, 5}
	fmt.Println("original arr : ", darr)
	max := 0
	sum := 0
	start := 0
	//ansstart:=-1
	ansend := -1
	for i := 0; i < len(darr); i++ {
		if sum == 0 {
			start = i
		}
		sum += darr[i]
		if sum > max {
			max = sum
			ansend = i
		}
		if sum < 0 {
			sum = 0
		}

	}
	fmt.Println("the maximam sub array from the original array : ", max)
	fmt.Println("this is the maximum sum of the arr : ", darr[start:ansend])
}
