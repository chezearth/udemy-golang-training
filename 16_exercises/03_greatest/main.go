package main

import "fmt"

func main() {
	n := greatest(3, -4, 10, 12, -7, 17, 1, -4)
	fmt.Println(n)
}

func greatest(list ...int) int {
	fmt.Printf("%T \n", list)
	var x int
	for _, v := range list {
		if x < v {
			x = v
		}
	}
	return x
}
