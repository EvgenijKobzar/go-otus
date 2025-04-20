package main

import (
	"fmt"
)

const NewLineChar = "\n"

func main() {

	str8х8 := " # # # #\n" +
		" # # # #\n" +
		" # # # #\n" +
		" # # # #\n"
	fmt.Println(str8х8)

	x := 16
	y := 16
	fmt.Println(render(x, y))
}

func render(x int, y int) string {
	result := ""

	for i := 0; i < y; i++ {
		result += axisX(x) + NewLineChar
	}
	return result
}

func axisX(len int) string {
	result := ""

	for i := 0; i < len; i++ {
		if i%2 == 0 {
			result += " "
		} else {
			result += "#"
		}
	}

	return result
}
