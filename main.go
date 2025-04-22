package main

import (
	"fmt"
	"strings"
)

const NewLineChar = "\n"

func main() {

	var x int
	var y int

	fmt.Println("Укажите колличество по горизонтали (X)")
	fmt.Scanf("%d\n", &x)

	fmt.Println("Укажите колличество по вертикали (Y)")
	fmt.Scanf("%d\n", &y)

	fmt.Println(render(x, y))
}

func render(x int, y int) string {
	result := strings.Builder{}

	for i := 0; i < y; i++ {

		if i%2 == 0 {
			result.WriteString(makeSequence(x, "#", " "))
		} else {
			result.WriteString(makeSequence(x, " ", "#"))
		}

		result.WriteString(NewLineChar)
	}
	return result.String()
}

func makeSequence(len int, symbol1 string, symbol2 string) string {
	result := strings.Builder{}

	for i := 0; i < len; i++ {
		if i%2 == 0 {
			result.WriteString(symbol1)
		} else {
			result.WriteString(symbol2)
		}
	}

	return result.String()
}
