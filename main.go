package main

import (
	"fmt"
)

const NewLineChar = "\n"

func main() {

	//str8х8 := " # # # #\n" +
	//	"# # # # \n" +
	//	" # # # #\n" +
	//	"# # # # \n"
	//fmt.Println(str8х8)

	x := 16
	y := 16
	fmt.Println(render(x, y))
}

func render(x int, y int) string {
	result := ""

	for i := 0; i < y; i++ {

		if i%2 == 0 {
			result += makeSequence(x, "#", " ")
		} else {
			result += makeSequence(x, " ", "#")
		}

		result += NewLineChar
	}
	return result
}

func makeSequence(len int, symbol1 string, symbol2 string) string {
	result := ""

	for i := 0; i < len; i++ {
		if i%2 == 0 {
			result += symbol1
		} else {
			result += symbol2
		}
	}

	return result
}
