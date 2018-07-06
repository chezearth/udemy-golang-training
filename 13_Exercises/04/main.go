package main

import (
	"fmt"
	"strconv"
)

func main() {
	var first, second, rem int
	var order string
	fmt.Print("Enter two numbers separated by a space: ")
	fmt.Scanln(&first, &second)
	if first > second && second != 0 {
		order = strconv.Itoa(first) + " / " + strconv.Itoa(second) + ": "
		rem = first % second
	} else if first != 0 {
		order = strconv.Itoa(second) + " / " + strconv.Itoa(first) + " = "
		rem = second % first
	} else {
		order = "No input:"
	}
	fmt.Println(order, rem)
}
