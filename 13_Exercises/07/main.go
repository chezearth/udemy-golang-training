package main

import "fmt"

var sum int

func main() {
	for i := 0; i < 1000; i++ {
		switch {
		case i%3 == 0, i%5 == 0:
			sum += i
			//			fmt.Printf("iter: %d, sum: %d\n", i, sum)
		default:
			//			fmt.Printf("iter: %d\n", i)
		}
	}
	fmt.Printf("The sum is %d\n", sum)
}
