package main

import "fmt"

func main() {

	half := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}

	var i int
	for i = 1; i <= 20; i++ {
		x, even := half(i)
		fmt.Println(x, even)
	}
}
